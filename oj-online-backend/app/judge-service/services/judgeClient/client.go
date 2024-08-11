package judgeClient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pengdahong1225/Oj-Online-Server/app/judge-service/services/goroutinePool"
	"github.com/pengdahong1225/Oj-Online-Server/app/judge-service/services/redis"
	"github.com/pengdahong1225/Oj-Online-Server/app/judge-service/setting"
	"github.com/pengdahong1225/Oj-Online-Server/app/question-service/models"
	"github.com/pengdahong1225/Oj-Online-Server/pkg/registry"
	"github.com/pengdahong1225/Oj-Online-Server/pkg/settings"
	pb "github.com/pengdahong1225/Oj-Online-Server/proto"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	baseUrl    string
	exeName    string            // 执行文件名
	runResults chan SubmitResult // 运行结果
)

func Init() error {
	srv, err := settings.GetSystemConf(setting.Instance().SystemConfigs, "judge-service")
	if err != nil {
		return err
	}

	baseUrl = fmt.Sprintf("http://%s:%d", srv.Host, srv.Port)
	runResults = make(chan SubmitResult, 256)
	exeName = "main"

	return nil
}

type Param struct {
	uid           int64
	problemID     int64
	compileConfig *ProblemConfig
	runConfig     *ProblemConfig
	content       string // 源代码
	testCases     []TestCase
	fileIds       map[string]string // 文件id
}

// Handle 判题服务入口
func Handle(uid int64, form *models.SubmitForm) []SubmitResult {
	// 退出之后，需要将状态置为UPStateExited，主要针对异常退出的情况，正常退出会设置状态
	defer func() {
		if err := redis.SetUPState(uid, form.ProblemID, int(pb.SubmitState_UPStateExited)); err != nil {
			logrus.Errorln(err.Error())
		}
	}()

	// 设置“提交”状态
	if err := redis.SetUPState(uid, form.ProblemID, int(pb.SubmitState_UPStateNormal)); err != nil {
		logrus.Errorln(err.Error())
		return nil
	}

	ok, param := preAction(uid, form)
	if !ok {
		logrus.Errorln("预处理失败")
		return nil
	}

	start := time.Now()
	res := doAction(param)
	duration := time.Now().Sub(start).Milliseconds()
	logrus.Infoln("---judgeClient.Handle--- uid:%d, problemID:%d, total-cost:%d ms", uid, form.ProblemID, duration)
	return res
}

func preAction(uid int64, form *models.SubmitForm) (bool, *Param) {
	param := &Param{}

	// 读取题目配置
	hotData := getProblemHotData(form.ProblemID)
	compileConfig := &ProblemConfig{}
	if err := json.Unmarshal([]byte(hotData.CompileConfig), compileConfig); err != nil {
		logrus.Errorln(err.Error())
		return false, nil
	}
	runConfig := &ProblemConfig{}
	if err := json.Unmarshal([]byte(hotData.RunConfig), runConfig); err != nil {
		logrus.Errorln(err.Error())
		return false, nil
	}
	var testCases []TestCase
	if err := json.Unmarshal([]byte(hotData.TestCase), &testCases); err != nil {
		logrus.Errorln(err.Error())
		return false, nil
	}

	param.uid = uid
	param.problemID = form.ProblemID
	param.compileConfig = compileConfig
	param.runConfig = runConfig
	param.content = form.Code
	param.testCases = testCases
	return true, param
}

// 操作(编译，运行，评判)，操作的上下文信息需要缓存到redis
// redis需要持久化的信息：
// 1.本次提交的状态
// 2.编译结果
// 3.运行结果
// 4.评判结果
func doAction(param *Param) []SubmitResult {
	handler := &Handler{}
	results := make([]SubmitResult, 0)
	// 设置题目状态[编译]
	if err := redis.SetUPState(param.uid, param.problemID, int(pb.SubmitState_UPStateCompiling)); err != nil {
		logrus.Errorln(err.Error())
	}
	compileResult, err := handler.compile(param)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil
	}
	if compileResult.Status != "Accepted" {
		compileResult.Content = "编译失败"
		results = append(results, *compileResult)
		// 更新状态
		if err := redis.SetUPState(param.uid, param.problemID, int(pb.SubmitState_UPStateExited)); err != nil {
			logrus.Errorln(err.Error())
			return nil
		}
		return results
	}
	compileResult.Content = "编译成功"
	results = append(results, *compileResult)

	// 编译成功 => 设置题目状态[判题中]
	// 运行和判题是协同的步骤，由两个协程去同时进行，通过channel通信
	param.fileIds = compileResult.FileIds
	if err := redis.SetUPState(param.uid, param.problemID, int(pb.SubmitState_UPStateJudging)); err != nil {
		logrus.Errorln(err.Error())
	}
	wgRun := new(sync.WaitGroup)
	wgRun.Add(1)
	goroutinePool.PoolInstance.Submit(func() {
		defer wgRun.Done()
		handler.run(param)
	})
	wgJudge := new(sync.WaitGroup)
	wgJudge.Add(1)
	goroutinePool.PoolInstance.Submit(func() {
		defer wgJudge.Done()
		judgeResults := handler.judge()
		results = append(results, judgeResults...)
	})

	wgRun.Wait()
	// 关闭管道，触发后续goroutine结束
	close(runResults)
	wgJudge.Wait()

	return results
}

// 获取题目热点数据
// cache中获取失败就去db获取
func getProblemHotData(ProblemID int64) *ProblemHotData {
	// cache
	data, err := redis.GetProblemHotData(ProblemID)
	if err != nil {
		logrus.Errorln(err.Error())
		return nil
	}
	if data == "" {
		// db
		dbConn, err := registry.NewDBConnection(setting.Instance().RegistryConfig)
		if err != nil {
			logrus.Errorf("db服连接失败:%s\n", err.Error())
			return nil
		}
		defer dbConn.Close()
		client := pb.NewDBServiceClient(dbConn)
		request := &pb.GetProblemHotDataRequest{
			ProblemId: ProblemID,
		}
		res, err := client.GetProblemHotData(context.Background(), request)
		if err != nil {
			logrus.Errorln(err.Error())
			return nil
		}
		data = res.Data
	}

	hotData := &ProblemHotData{}
	if err := json.Unmarshal([]byte(data), hotData); err != nil {
		logrus.Errorln(err.Error())
		return nil
	}
	return hotData
}
