// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: register.proto

package pb

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

// RegiterClient is the client API for Regiter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegiterClient interface {
	SendSMS(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSMSResponse, error)
	VerifySMS(ctx context.Context, in *VerifySMSRequest, opts ...grpc.CallOption) (*VerifySMSResponse, error)
}

type regiterClient struct {
	cc grpc.ClientConnInterface
}

func NewRegiterClient(cc grpc.ClientConnInterface) RegiterClient {
	return &regiterClient{cc}
}

func (c *regiterClient) SendSMS(ctx context.Context, in *SendSMSRequest, opts ...grpc.CallOption) (*SendSMSResponse, error) {
	out := new(SendSMSResponse)
	err := c.cc.Invoke(ctx, "/pb.Regiter/SendSMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regiterClient) VerifySMS(ctx context.Context, in *VerifySMSRequest, opts ...grpc.CallOption) (*VerifySMSResponse, error) {
	out := new(VerifySMSResponse)
	err := c.cc.Invoke(ctx, "/pb.Regiter/VerifySMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegiterServer is the server API for Regiter service.
// All implementations must embed UnimplementedRegiterServer
// for forward compatibility
type RegiterServer interface {
	SendSMS(context.Context, *SendSMSRequest) (*SendSMSResponse, error)
	VerifySMS(context.Context, *VerifySMSRequest) (*VerifySMSResponse, error)
	mustEmbedUnimplementedRegiterServer()
}

// UnimplementedRegiterServer must be embedded to have forward compatible implementations.
type UnimplementedRegiterServer struct {
}

func (UnimplementedRegiterServer) SendSMS(context.Context, *SendSMSRequest) (*SendSMSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMS not implemented")
}
func (UnimplementedRegiterServer) VerifySMS(context.Context, *VerifySMSRequest) (*VerifySMSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySMS not implemented")
}
func (UnimplementedRegiterServer) mustEmbedUnimplementedRegiterServer() {}

// UnsafeRegiterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegiterServer will
// result in compilation errors.
type UnsafeRegiterServer interface {
	mustEmbedUnimplementedRegiterServer()
}

func RegisterRegiterServer(s grpc.ServiceRegistrar, srv RegiterServer) {
	s.RegisterService(&Regiter_ServiceDesc, srv)
}

func _Regiter_SendSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegiterServer).SendSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Regiter/SendSMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegiterServer).SendSMS(ctx, req.(*SendSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Regiter_VerifySMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifySMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegiterServer).VerifySMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Regiter/VerifySMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegiterServer).VerifySMS(ctx, req.(*VerifySMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Regiter_ServiceDesc is the grpc.ServiceDesc for Regiter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Regiter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Regiter",
	HandlerType: (*RegiterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSMS",
			Handler:    _Regiter_SendSMS_Handler,
		},
		{
			MethodName: "VerifySMS",
			Handler:    _Regiter_VerifySMS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}
