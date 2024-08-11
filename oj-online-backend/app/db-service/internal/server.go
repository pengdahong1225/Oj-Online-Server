package internal

import (
	"fmt"
	"github.com/pengdahong1225/Oj-Online-Server/app/db-service/internal/daemon"
	"github.com/pengdahong1225/Oj-Online-Server/app/db-service/internal/handler"
	goroutinePool "github.com/pengdahong1225/Oj-Online-Server/app/db-service/services/goroutinePoll"
	"github.com/pengdahong1225/Oj-Online-Server/app/db-service/setting"
	"github.com/pengdahong1225/Oj-Online-Server/pkg/registry"
	"github.com/pengdahong1225/Oj-Online-Server/pkg/settings"
	"github.com/pengdahong1225/Oj-Online-Server/pkg/utils"
	"github.com/pengdahong1225/Oj-Online-Server/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"sync"
)

type Server struct {
}

func (receiver Server) Start() {
	// 后台-排行榜服务
	wg := new(sync.WaitGroup)
	wg.Add(2)
	err := goroutinePool.Instance().Submit(func() {
		defer wg.Done()
		daemon.StartDaemon()
	})
	if err != nil {
		panic(err)
	}
	// DB服务
	err = goroutinePool.Instance().Submit(func() {
		defer wg.Done()
		StartRPCServer()
	})
	if err != nil {
		panic(err)
	}
	wg.Wait()
}

func StartRPCServer() {
	system, err := settings.GetSystemConf(setting.Instance().SystemConfigs, "db-service")
	if err != nil {
		panic(err)
	}
	// 获取ip地址
	ip, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	// 监听端口
	netAddr := fmt.Sprintf("%s:%d", ip.String(), system.Port)
	listener, err := net.Listen("tcp", netAddr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	// 健康检查
	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthcheck)
	// 注册
	register, _ := registry.NewRegistry(setting.Instance().RegistryConfig)
	if err := register.RegisterService(system.Name, ip.String(), system.Port); err != nil {
		panic(err)
	}

	// 注册并启动db服务
	dbSrv := handler.DBServiceServer{}
	pb.RegisterDBServiceServer(grpcServer, &dbSrv)
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
