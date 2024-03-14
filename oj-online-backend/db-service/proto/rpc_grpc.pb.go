// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.10.0
// source: rpc.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DBService_GetUserData_FullMethodName            = "/DBService/GetUserData"
	DBService_CreateUserData_FullMethodName         = "/DBService/CreateUserData"
	DBService_UpdateUserData_FullMethodName         = "/DBService/UpdateUserData"
	DBService_DeleteUserData_FullMethodName         = "/DBService/DeleteUserData"
	DBService_GetUserList_FullMethodName            = "/DBService/GetUserList"
	DBService_GetQuestionData_FullMethodName        = "/DBService/GetQuestionData"
	DBService_CreateQuestionData_FullMethodName     = "/DBService/CreateQuestionData"
	DBService_UpdateQuestionData_FullMethodName     = "/DBService/UpdateQuestionData"
	DBService_DeleteQuestionData_FullMethodName     = "/DBService/DeleteQuestionData"
	DBService_GetQuestionList_FullMethodName        = "/DBService/GetQuestionList"
	DBService_QueryQuestionWithName_FullMethodName  = "/DBService/QueryQuestionWithName"
	DBService_GetUserSubmitRecord_FullMethodName    = "/DBService/GetUserSubmitRecord"
	DBService_UpdateUserSubmitRecord_FullMethodName = "/DBService/UpdateUserSubmitRecord"
)

// DBServiceClient is the client API for DBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DBServiceClient interface {
	GetUserData(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateUserData(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUserData(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUserData(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error)
	GetQuestionData(ctx context.Context, in *GetQuestionRequest, opts ...grpc.CallOption) (*GetQuestionResponse, error)
	CreateQuestionData(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionResponse, error)
	UpdateQuestionData(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteQuestionData(ctx context.Context, in *DeleteQuestionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetQuestionList(ctx context.Context, in *GetQuestionListRequest, opts ...grpc.CallOption) (*GetQuestionListResponse, error)
	QueryQuestionWithName(ctx context.Context, in *QueryQuestionWithNameRequest, opts ...grpc.CallOption) (*QueryQuestionWithNameResponse, error)
	GetUserSubmitRecord(ctx context.Context, in *GetUserSubmitRecordRequest, opts ...grpc.CallOption) (*GetUserSubmitRecordResponse, error)
	UpdateUserSubmitRecord(ctx context.Context, in *UpdateUserSubmitRecordRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type dBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDBServiceClient(cc grpc.ClientConnInterface) DBServiceClient {
	return &dBServiceClient{cc}
}

func (c *dBServiceClient) GetUserData(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, DBService_GetUserData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) CreateUserData(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, DBService_CreateUserData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) UpdateUserData(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DBService_UpdateUserData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) DeleteUserData(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DBService_DeleteUserData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListResponse, error) {
	out := new(GetUserListResponse)
	err := c.cc.Invoke(ctx, DBService_GetUserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) GetQuestionData(ctx context.Context, in *GetQuestionRequest, opts ...grpc.CallOption) (*GetQuestionResponse, error) {
	out := new(GetQuestionResponse)
	err := c.cc.Invoke(ctx, DBService_GetQuestionData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) CreateQuestionData(ctx context.Context, in *CreateQuestionRequest, opts ...grpc.CallOption) (*CreateQuestionResponse, error) {
	out := new(CreateQuestionResponse)
	err := c.cc.Invoke(ctx, DBService_CreateQuestionData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) UpdateQuestionData(ctx context.Context, in *UpdateQuestionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DBService_UpdateQuestionData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) DeleteQuestionData(ctx context.Context, in *DeleteQuestionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DBService_DeleteQuestionData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) GetQuestionList(ctx context.Context, in *GetQuestionListRequest, opts ...grpc.CallOption) (*GetQuestionListResponse, error) {
	out := new(GetQuestionListResponse)
	err := c.cc.Invoke(ctx, DBService_GetQuestionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) QueryQuestionWithName(ctx context.Context, in *QueryQuestionWithNameRequest, opts ...grpc.CallOption) (*QueryQuestionWithNameResponse, error) {
	out := new(QueryQuestionWithNameResponse)
	err := c.cc.Invoke(ctx, DBService_QueryQuestionWithName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) GetUserSubmitRecord(ctx context.Context, in *GetUserSubmitRecordRequest, opts ...grpc.CallOption) (*GetUserSubmitRecordResponse, error) {
	out := new(GetUserSubmitRecordResponse)
	err := c.cc.Invoke(ctx, DBService_GetUserSubmitRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dBServiceClient) UpdateUserSubmitRecord(ctx context.Context, in *UpdateUserSubmitRecordRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, DBService_UpdateUserSubmitRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DBServiceServer is the server API for DBService service.
// All implementations must embed UnimplementedDBServiceServer
// for forward compatibility
type DBServiceServer interface {
	GetUserData(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateUserData(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUserData(context.Context, *UpdateUserRequest) (*empty.Empty, error)
	DeleteUserData(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	GetQuestionData(context.Context, *GetQuestionRequest) (*GetQuestionResponse, error)
	CreateQuestionData(context.Context, *CreateQuestionRequest) (*CreateQuestionResponse, error)
	UpdateQuestionData(context.Context, *UpdateQuestionRequest) (*empty.Empty, error)
	DeleteQuestionData(context.Context, *DeleteQuestionRequest) (*empty.Empty, error)
	GetQuestionList(context.Context, *GetQuestionListRequest) (*GetQuestionListResponse, error)
	QueryQuestionWithName(context.Context, *QueryQuestionWithNameRequest) (*QueryQuestionWithNameResponse, error)
	GetUserSubmitRecord(context.Context, *GetUserSubmitRecordRequest) (*GetUserSubmitRecordResponse, error)
	UpdateUserSubmitRecord(context.Context, *UpdateUserSubmitRecordRequest) (*empty.Empty, error)
	mustEmbedUnimplementedDBServiceServer()
}

// UnimplementedDBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDBServiceServer struct {
}

func (UnimplementedDBServiceServer) GetUserData(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserData not implemented")
}
func (UnimplementedDBServiceServer) CreateUserData(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserData not implemented")
}
func (UnimplementedDBServiceServer) UpdateUserData(context.Context, *UpdateUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserData not implemented")
}
func (UnimplementedDBServiceServer) DeleteUserData(context.Context, *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserData not implemented")
}
func (UnimplementedDBServiceServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedDBServiceServer) GetQuestionData(context.Context, *GetQuestionRequest) (*GetQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionData not implemented")
}
func (UnimplementedDBServiceServer) CreateQuestionData(context.Context, *CreateQuestionRequest) (*CreateQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuestionData not implemented")
}
func (UnimplementedDBServiceServer) UpdateQuestionData(context.Context, *UpdateQuestionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuestionData not implemented")
}
func (UnimplementedDBServiceServer) DeleteQuestionData(context.Context, *DeleteQuestionRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuestionData not implemented")
}
func (UnimplementedDBServiceServer) GetQuestionList(context.Context, *GetQuestionListRequest) (*GetQuestionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestionList not implemented")
}
func (UnimplementedDBServiceServer) QueryQuestionWithName(context.Context, *QueryQuestionWithNameRequest) (*QueryQuestionWithNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryQuestionWithName not implemented")
}
func (UnimplementedDBServiceServer) GetUserSubmitRecord(context.Context, *GetUserSubmitRecordRequest) (*GetUserSubmitRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSubmitRecord not implemented")
}
func (UnimplementedDBServiceServer) UpdateUserSubmitRecord(context.Context, *UpdateUserSubmitRecordRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserSubmitRecord not implemented")
}
func (UnimplementedDBServiceServer) mustEmbedUnimplementedDBServiceServer() {}

// UnsafeDBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DBServiceServer will
// result in compilation errors.
type UnsafeDBServiceServer interface {
	mustEmbedUnimplementedDBServiceServer()
}

func RegisterDBServiceServer(s grpc.ServiceRegistrar, srv DBServiceServer) {
	s.RegisterService(&DBService_ServiceDesc, srv)
}

func _DBService_GetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_GetUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetUserData(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_CreateUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).CreateUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_CreateUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).CreateUserData(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_UpdateUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).UpdateUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_UpdateUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).UpdateUserData(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_DeleteUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).DeleteUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_DeleteUserData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).DeleteUserData(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_GetUserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_GetQuestionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetQuestionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_GetQuestionData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetQuestionData(ctx, req.(*GetQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_CreateQuestionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).CreateQuestionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_CreateQuestionData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).CreateQuestionData(ctx, req.(*CreateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_UpdateQuestionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).UpdateQuestionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_UpdateQuestionData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).UpdateQuestionData(ctx, req.(*UpdateQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_DeleteQuestionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).DeleteQuestionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_DeleteQuestionData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).DeleteQuestionData(ctx, req.(*DeleteQuestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_GetQuestionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuestionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetQuestionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_GetQuestionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetQuestionList(ctx, req.(*GetQuestionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_QueryQuestionWithName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryQuestionWithNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).QueryQuestionWithName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_QueryQuestionWithName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).QueryQuestionWithName(ctx, req.(*QueryQuestionWithNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_GetUserSubmitRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSubmitRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).GetUserSubmitRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_GetUserSubmitRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).GetUserSubmitRecord(ctx, req.(*GetUserSubmitRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DBService_UpdateUserSubmitRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserSubmitRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DBServiceServer).UpdateUserSubmitRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DBService_UpdateUserSubmitRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DBServiceServer).UpdateUserSubmitRecord(ctx, req.(*UpdateUserSubmitRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DBService_ServiceDesc is the grpc.ServiceDesc for DBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DBService",
	HandlerType: (*DBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserData",
			Handler:    _DBService_GetUserData_Handler,
		},
		{
			MethodName: "CreateUserData",
			Handler:    _DBService_CreateUserData_Handler,
		},
		{
			MethodName: "UpdateUserData",
			Handler:    _DBService_UpdateUserData_Handler,
		},
		{
			MethodName: "DeleteUserData",
			Handler:    _DBService_DeleteUserData_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _DBService_GetUserList_Handler,
		},
		{
			MethodName: "GetQuestionData",
			Handler:    _DBService_GetQuestionData_Handler,
		},
		{
			MethodName: "CreateQuestionData",
			Handler:    _DBService_CreateQuestionData_Handler,
		},
		{
			MethodName: "UpdateQuestionData",
			Handler:    _DBService_UpdateQuestionData_Handler,
		},
		{
			MethodName: "DeleteQuestionData",
			Handler:    _DBService_DeleteQuestionData_Handler,
		},
		{
			MethodName: "GetQuestionList",
			Handler:    _DBService_GetQuestionList_Handler,
		},
		{
			MethodName: "QueryQuestionWithName",
			Handler:    _DBService_QueryQuestionWithName_Handler,
		},
		{
			MethodName: "GetUserSubmitRecord",
			Handler:    _DBService_GetUserSubmitRecord_Handler,
		},
		{
			MethodName: "UpdateUserSubmitRecord",
			Handler:    _DBService_UpdateUserSubmitRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
