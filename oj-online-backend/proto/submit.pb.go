// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.0
// source: submit.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 用户状态枚举，如果没有查询到，意味着用户最近没有提交过题目
type UserState int32

const (
	UserState_UserStateNormal  UserState = 0
	UserState_UserStateJudging UserState = 1
)

// Enum value maps for UserState.
var (
	UserState_name = map[int32]string{
		0: "UserStateNormal",
		1: "UserStateJudging",
	}
	UserState_value = map[string]int32{
		"UserStateNormal":  0,
		"UserStateJudging": 1,
	}
)

func (x UserState) Enum() *UserState {
	p := new(UserState)
	*p = x
	return p
}

func (x UserState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserState) Descriptor() protoreflect.EnumDescriptor {
	return file_submit_proto_enumTypes[0].Descriptor()
}

func (UserState) Type() protoreflect.EnumType {
	return &file_submit_proto_enumTypes[0]
}

func (x UserState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserState.Descriptor instead.
func (UserState) EnumDescriptor() ([]byte, []int) {
	return file_submit_proto_rawDescGZIP(), []int{0}
}

// “提交”状态枚举，如果没有查询到状态，就意味着最近没有提交题目or题目提交过期了
type SubmitState int32

const (
	SubmitState_UPStateNormal    SubmitState = 0 // 初始状态
	SubmitState_UPStateCompiling SubmitState = 1 // 正在编译
	SubmitState_UPStateJudging   SubmitState = 2 // 已编译成功，正在判题
	SubmitState_UPStateExited    SubmitState = 3 // 已退出 -> 如何查询到这个状态，就意味着可以查询结果了
)

// Enum value maps for SubmitState.
var (
	SubmitState_name = map[int32]string{
		0: "UPStateNormal",
		1: "UPStateCompiling",
		2: "UPStateJudging",
		3: "UPStateExited",
	}
	SubmitState_value = map[string]int32{
		"UPStateNormal":    0,
		"UPStateCompiling": 1,
		"UPStateJudging":   2,
		"UPStateExited":    3,
	}
)

func (x SubmitState) Enum() *SubmitState {
	p := new(SubmitState)
	*p = x
	return p
}

func (x SubmitState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubmitState) Descriptor() protoreflect.EnumDescriptor {
	return file_submit_proto_enumTypes[1].Descriptor()
}

func (SubmitState) Type() protoreflect.EnumType {
	return &file_submit_proto_enumTypes[1]
}

func (x SubmitState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubmitState.Descriptor instead.
func (SubmitState) EnumDescriptor() ([]byte, []int) {
	return file_submit_proto_rawDescGZIP(), []int{1}
}

type SubmitForm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId int64  `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Lang      string `protobuf:"bytes,3,opt,name=lang,proto3" json:"lang,omitempty"`
	Code      string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SubmitForm) Reset() {
	*x = SubmitForm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitForm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitForm) ProtoMessage() {}

func (x *SubmitForm) ProtoReflect() protoreflect.Message {
	mi := &file_submit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitForm.ProtoReflect.Descriptor instead.
func (*SubmitForm) Descriptor() ([]byte, []int) {
	return file_submit_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitForm) GetProblemId() int64 {
	if x != nil {
		return x.ProblemId
	}
	return 0
}

func (x *SubmitForm) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SubmitForm) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *SubmitForm) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_submit_proto protoreflect.FileDescriptor

var file_submit_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69,
	0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x2a, 0x36, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x75, 0x64, 0x67, 0x69, 0x6e, 0x67, 0x10,
	0x01, 0x2a, 0x5d, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x11, 0x0a, 0x0d, 0x55, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x50, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x4a, 0x75, 0x64, 0x67, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x50, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x69, 0x74, 0x65, 0x64, 0x10, 0x03,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_submit_proto_rawDescOnce sync.Once
	file_submit_proto_rawDescData = file_submit_proto_rawDesc
)

func file_submit_proto_rawDescGZIP() []byte {
	file_submit_proto_rawDescOnce.Do(func() {
		file_submit_proto_rawDescData = protoimpl.X.CompressGZIP(file_submit_proto_rawDescData)
	})
	return file_submit_proto_rawDescData
}

var file_submit_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_submit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_submit_proto_goTypes = []any{
	(UserState)(0),     // 0: UserState
	(SubmitState)(0),   // 1: SubmitState
	(*SubmitForm)(nil), // 2: SubmitForm
}
var file_submit_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_submit_proto_init() }
func file_submit_proto_init() {
	if File_submit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submit_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SubmitForm); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_submit_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_submit_proto_goTypes,
		DependencyIndexes: file_submit_proto_depIdxs,
		EnumInfos:         file_submit_proto_enumTypes,
		MessageInfos:      file_submit_proto_msgTypes,
	}.Build()
	File_submit_proto = out.File
	file_submit_proto_rawDesc = nil
	file_submit_proto_goTypes = nil
	file_submit_proto_depIdxs = nil
}
