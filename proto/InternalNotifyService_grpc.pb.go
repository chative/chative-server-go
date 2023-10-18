// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: InternalNotifyService.proto

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

// InternalNotifyServiceClient is the client API for InternalNotifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalNotifyServiceClient interface {
	SendNotify(ctx context.Context, in *NotifySendRequest, opts ...grpc.CallOption) (*BaseResponse, error)
}

type internalNotifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalNotifyServiceClient(cc grpc.ClientConnInterface) InternalNotifyServiceClient {
	return &internalNotifyServiceClient{cc}
}

func (c *internalNotifyServiceClient) SendNotify(ctx context.Context, in *NotifySendRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalNotifyService/sendNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalNotifyServiceServer is the server API for InternalNotifyService service.
// All implementations must embed UnimplementedInternalNotifyServiceServer
// for forward compatibility
type InternalNotifyServiceServer interface {
	SendNotify(context.Context, *NotifySendRequest) (*BaseResponse, error)
	mustEmbedUnimplementedInternalNotifyServiceServer()
}

// UnimplementedInternalNotifyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInternalNotifyServiceServer struct {
}

func (UnimplementedInternalNotifyServiceServer) SendNotify(context.Context, *NotifySendRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotify not implemented")
}
func (UnimplementedInternalNotifyServiceServer) mustEmbedUnimplementedInternalNotifyServiceServer() {}

// UnsafeInternalNotifyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalNotifyServiceServer will
// result in compilation errors.
type UnsafeInternalNotifyServiceServer interface {
	mustEmbedUnimplementedInternalNotifyServiceServer()
}

func RegisterInternalNotifyServiceServer(s grpc.ServiceRegistrar, srv InternalNotifyServiceServer) {
	s.RegisterService(&InternalNotifyService_ServiceDesc, srv)
}

func _InternalNotifyService_SendNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifySendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalNotifyServiceServer).SendNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalNotifyService/sendNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalNotifyServiceServer).SendNotify(ctx, req.(*NotifySendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalNotifyService_ServiceDesc is the grpc.ServiceDesc for InternalNotifyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalNotifyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InternalNotifyService",
	HandlerType: (*InternalNotifyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendNotify",
			Handler:    _InternalNotifyService_SendNotify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "InternalNotifyService.proto",
}