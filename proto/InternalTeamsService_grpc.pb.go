// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: InternalTeamsService.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InternalTeamsServiceClient is the client API for InternalTeamsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalTeamsServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *Step, opts ...grpc.CallOption) (*GetResponse, error)
	Join(ctx context.Context, in *JoinLeaveRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	Leave(ctx context.Context, in *JoinLeaveRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	GetTree(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*GetResponse, error)
	CreateOrUpdate(ctx context.Context, in *CreateOrUpdateRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	Delete(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type internalTeamsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalTeamsServiceClient(cc grpc.ClientConnInterface) InternalTeamsServiceClient {
	return &internalTeamsServiceClient{cc}
}

func (c *internalTeamsServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/InternalTeamsService/get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalTeamsServiceClient) GetAll(ctx context.Context, in *Step, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/InternalTeamsService/getAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalTeamsServiceClient) Join(ctx context.Context, in *JoinLeaveRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalTeamsService/join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalTeamsServiceClient) Leave(ctx context.Context, in *JoinLeaveRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalTeamsService/leave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalTeamsServiceClient) GetTree(ctx context.Context, in *GetTreeRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/InternalTeamsService/getTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalTeamsServiceClient) CreateOrUpdate(ctx context.Context, in *CreateOrUpdateRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalTeamsService/createOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalTeamsServiceClient) Delete(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalTeamsService/delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalTeamsServiceServer is the server API for InternalTeamsService service.
// All implementations must embed UnimplementedInternalTeamsServiceServer
// for forward compatibility
type InternalTeamsServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *Step) (*GetResponse, error)
	Join(context.Context, *JoinLeaveRequest) (*BaseResponse, error)
	Leave(context.Context, *JoinLeaveRequest) (*BaseResponse, error)
	GetTree(context.Context, *GetTreeRequest) (*GetResponse, error)
	CreateOrUpdate(context.Context, *CreateOrUpdateRequest) (*BaseResponse, error)
	Delete(context.Context, *GetRequest) (*BaseResponse, error)
	mustEmbedUnimplementedInternalTeamsServiceServer()
}

// UnimplementedInternalTeamsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInternalTeamsServiceServer struct {
}

func (UnimplementedInternalTeamsServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedInternalTeamsServiceServer) GetAll(context.Context, *Step) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedInternalTeamsServiceServer) Join(context.Context, *JoinLeaveRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedInternalTeamsServiceServer) Leave(context.Context, *JoinLeaveRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (UnimplementedInternalTeamsServiceServer) GetTree(context.Context, *GetTreeRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTree not implemented")
}
func (UnimplementedInternalTeamsServiceServer) CreateOrUpdate(context.Context, *CreateOrUpdateRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (UnimplementedInternalTeamsServiceServer) Delete(context.Context, *GetRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedInternalTeamsServiceServer) mustEmbedUnimplementedInternalTeamsServiceServer() {}

// UnsafeInternalTeamsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalTeamsServiceServer will
// result in compilation errors.
type UnsafeInternalTeamsServiceServer interface {
	mustEmbedUnimplementedInternalTeamsServiceServer()
}

func RegisterInternalTeamsServiceServer(s grpc.ServiceRegistrar, srv InternalTeamsServiceServer) {
	s.RegisterService(&InternalTeamsService_ServiceDesc, srv)
}

func _InternalTeamsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalTeamsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalTeamsService/get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalTeamsServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalTeamsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Step)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalTeamsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalTeamsService/getAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalTeamsServiceServer).GetAll(ctx, req.(*Step))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalTeamsService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalTeamsServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalTeamsService/join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalTeamsServiceServer).Join(ctx, req.(*JoinLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalTeamsService_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalTeamsServiceServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalTeamsService/leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalTeamsServiceServer).Leave(ctx, req.(*JoinLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalTeamsService_GetTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalTeamsServiceServer).GetTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalTeamsService/getTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalTeamsServiceServer).GetTree(ctx, req.(*GetTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalTeamsService_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalTeamsServiceServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalTeamsService/createOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalTeamsServiceServer).CreateOrUpdate(ctx, req.(*CreateOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalTeamsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalTeamsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalTeamsService/delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalTeamsServiceServer).Delete(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalTeamsService_ServiceDesc is the grpc.ServiceDesc for InternalTeamsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalTeamsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InternalTeamsService",
	HandlerType: (*InternalTeamsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "get",
			Handler:    _InternalTeamsService_Get_Handler,
		},
		{
			MethodName: "getAll",
			Handler:    _InternalTeamsService_GetAll_Handler,
		},
		{
			MethodName: "join",
			Handler:    _InternalTeamsService_Join_Handler,
		},
		{
			MethodName: "leave",
			Handler:    _InternalTeamsService_Leave_Handler,
		},
		{
			MethodName: "getTree",
			Handler:    _InternalTeamsService_GetTree_Handler,
		},
		{
			MethodName: "createOrUpdate",
			Handler:    _InternalTeamsService_CreateOrUpdate_Handler,
		},
		{
			MethodName: "delete",
			Handler:    _InternalTeamsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "InternalTeamsService.proto",
}
