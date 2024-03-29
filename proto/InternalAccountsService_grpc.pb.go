// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: InternalAccountsService.proto

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

// InternalAccountsServiceClient is the client API for InternalAccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InternalAccountsServiceClient interface {
	GetInfo(ctx context.Context, in *UidsRequest, opts ...grpc.CallOption) (*AccountInfoResponse, error)
	GetInfoByEmail(ctx context.Context, in *EmailsRequest, opts ...grpc.CallOption) (*AccountInfoResponse, error)
	Disable(ctx context.Context, in *UidsRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	Enable(ctx context.Context, in *UidsRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	GetAll(ctx context.Context, in *Step, opts ...grpc.CallOption) (*AccountInfoResponse, error)
	Edit(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	Renew(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	CreateAccount(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	QueryAccount(ctx context.Context, in *AccountQueryRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	KickOffDevice(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	GetUserTeams(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	SyncAccountBuInfo(ctx context.Context, in *SyncAccountBuRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	Inactive(ctx context.Context, in *UidsRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	DownloadAvatar(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	GenLoginInfo(ctx context.Context, in *LoginInfoReq, opts ...grpc.CallOption) (*BaseResponse, error)
	BlockConversation(ctx context.Context, in *BlockConversationRequest, opts ...grpc.CallOption) (*BaseResponse, error)
	GetConversationBlockStatus(ctx context.Context, in *GetConversationBlockReq, opts ...grpc.CallOption) (*BaseResponse, error)
}

type internalAccountsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInternalAccountsServiceClient(cc grpc.ClientConnInterface) InternalAccountsServiceClient {
	return &internalAccountsServiceClient{cc}
}

func (c *internalAccountsServiceClient) GetInfo(ctx context.Context, in *UidsRequest, opts ...grpc.CallOption) (*AccountInfoResponse, error) {
	out := new(AccountInfoResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/getInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) GetInfoByEmail(ctx context.Context, in *EmailsRequest, opts ...grpc.CallOption) (*AccountInfoResponse, error) {
	out := new(AccountInfoResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/getInfoByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) Disable(ctx context.Context, in *UidsRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/disable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) Enable(ctx context.Context, in *UidsRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/enable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) GetAll(ctx context.Context, in *Step, opts ...grpc.CallOption) (*AccountInfoResponse, error) {
	out := new(AccountInfoResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/getAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) Edit(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/edit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) Renew(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/renew", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) CreateAccount(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/createAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) QueryAccount(ctx context.Context, in *AccountQueryRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/queryAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) Upload(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/uploadAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) KickOffDevice(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/kickOffDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) GetUserTeams(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/getUserTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) SyncAccountBuInfo(ctx context.Context, in *SyncAccountBuRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/syncAccountBuInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) Inactive(ctx context.Context, in *UidsRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/inactive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) DownloadAvatar(ctx context.Context, in *AccountInfoRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/downloadAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) GenLoginInfo(ctx context.Context, in *LoginInfoReq, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/genLoginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) BlockConversation(ctx context.Context, in *BlockConversationRequest, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/blockConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalAccountsServiceClient) GetConversationBlockStatus(ctx context.Context, in *GetConversationBlockReq, opts ...grpc.CallOption) (*BaseResponse, error) {
	out := new(BaseResponse)
	err := c.cc.Invoke(ctx, "/InternalAccountsService/getConversationBlockStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalAccountsServiceServer is the server API for InternalAccountsService service.
// All implementations must embed UnimplementedInternalAccountsServiceServer
// for forward compatibility
type InternalAccountsServiceServer interface {
	GetInfo(context.Context, *UidsRequest) (*AccountInfoResponse, error)
	GetInfoByEmail(context.Context, *EmailsRequest) (*AccountInfoResponse, error)
	Disable(context.Context, *UidsRequest) (*BaseResponse, error)
	Enable(context.Context, *UidsRequest) (*BaseResponse, error)
	GetAll(context.Context, *Step) (*AccountInfoResponse, error)
	Edit(context.Context, *AccountInfoRequest) (*BaseResponse, error)
	Renew(context.Context, *AccountInfoRequest) (*BaseResponse, error)
	CreateAccount(context.Context, *AccountCreateRequest) (*BaseResponse, error)
	QueryAccount(context.Context, *AccountQueryRequest) (*BaseResponse, error)
	Upload(context.Context, *UploadRequest) (*BaseResponse, error)
	UploadAvatar(context.Context, *UploadAvatarRequest) (*BaseResponse, error)
	KickOffDevice(context.Context, *AccountInfoRequest) (*BaseResponse, error)
	GetUserTeams(context.Context, *TeamRequest) (*BaseResponse, error)
	SyncAccountBuInfo(context.Context, *SyncAccountBuRequest) (*BaseResponse, error)
	Inactive(context.Context, *UidsRequest) (*BaseResponse, error)
	DownloadAvatar(context.Context, *AccountInfoRequest) (*BaseResponse, error)
	GenLoginInfo(context.Context, *LoginInfoReq) (*BaseResponse, error)
	BlockConversation(context.Context, *BlockConversationRequest) (*BaseResponse, error)
	GetConversationBlockStatus(context.Context, *GetConversationBlockReq) (*BaseResponse, error)
	mustEmbedUnimplementedInternalAccountsServiceServer()
}

// UnimplementedInternalAccountsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInternalAccountsServiceServer struct {
}

func (UnimplementedInternalAccountsServiceServer) GetInfo(context.Context, *UidsRequest) (*AccountInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedInternalAccountsServiceServer) GetInfoByEmail(context.Context, *EmailsRequest) (*AccountInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfoByEmail not implemented")
}
func (UnimplementedInternalAccountsServiceServer) Disable(context.Context, *UidsRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedInternalAccountsServiceServer) Enable(context.Context, *UidsRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedInternalAccountsServiceServer) GetAll(context.Context, *Step) (*AccountInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedInternalAccountsServiceServer) Edit(context.Context, *AccountInfoRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edit not implemented")
}
func (UnimplementedInternalAccountsServiceServer) Renew(context.Context, *AccountInfoRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Renew not implemented")
}
func (UnimplementedInternalAccountsServiceServer) CreateAccount(context.Context, *AccountCreateRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedInternalAccountsServiceServer) QueryAccount(context.Context, *AccountQueryRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAccount not implemented")
}
func (UnimplementedInternalAccountsServiceServer) Upload(context.Context, *UploadRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedInternalAccountsServiceServer) UploadAvatar(context.Context, *UploadAvatarRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedInternalAccountsServiceServer) KickOffDevice(context.Context, *AccountInfoRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickOffDevice not implemented")
}
func (UnimplementedInternalAccountsServiceServer) GetUserTeams(context.Context, *TeamRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTeams not implemented")
}
func (UnimplementedInternalAccountsServiceServer) SyncAccountBuInfo(context.Context, *SyncAccountBuRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncAccountBuInfo not implemented")
}
func (UnimplementedInternalAccountsServiceServer) Inactive(context.Context, *UidsRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Inactive not implemented")
}
func (UnimplementedInternalAccountsServiceServer) DownloadAvatar(context.Context, *AccountInfoRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadAvatar not implemented")
}
func (UnimplementedInternalAccountsServiceServer) GenLoginInfo(context.Context, *LoginInfoReq) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenLoginInfo not implemented")
}
func (UnimplementedInternalAccountsServiceServer) BlockConversation(context.Context, *BlockConversationRequest) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockConversation not implemented")
}
func (UnimplementedInternalAccountsServiceServer) GetConversationBlockStatus(context.Context, *GetConversationBlockReq) (*BaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationBlockStatus not implemented")
}
func (UnimplementedInternalAccountsServiceServer) mustEmbedUnimplementedInternalAccountsServiceServer() {
}

// UnsafeInternalAccountsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InternalAccountsServiceServer will
// result in compilation errors.
type UnsafeInternalAccountsServiceServer interface {
	mustEmbedUnimplementedInternalAccountsServiceServer()
}

func RegisterInternalAccountsServiceServer(s grpc.ServiceRegistrar, srv InternalAccountsServiceServer) {
	s.RegisterService(&InternalAccountsService_ServiceDesc, srv)
}

func _InternalAccountsService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/getInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).GetInfo(ctx, req.(*UidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_GetInfoByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).GetInfoByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/getInfoByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).GetInfoByEmail(ctx, req.(*EmailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/disable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).Disable(ctx, req.(*UidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/enable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).Enable(ctx, req.(*UidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Step)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/getAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).GetAll(ctx, req.(*Step))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_Edit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).Edit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/edit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).Edit(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).Renew(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/createAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).CreateAccount(ctx, req.(*AccountCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_QueryAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).QueryAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/queryAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).QueryAccount(ctx, req.(*AccountQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).Upload(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_UploadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).UploadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/uploadAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).UploadAvatar(ctx, req.(*UploadAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_KickOffDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).KickOffDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/kickOffDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).KickOffDevice(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_GetUserTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).GetUserTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/getUserTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).GetUserTeams(ctx, req.(*TeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_SyncAccountBuInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncAccountBuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).SyncAccountBuInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/syncAccountBuInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).SyncAccountBuInfo(ctx, req.(*SyncAccountBuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_Inactive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UidsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).Inactive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/inactive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).Inactive(ctx, req.(*UidsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_DownloadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).DownloadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/downloadAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).DownloadAvatar(ctx, req.(*AccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_GenLoginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).GenLoginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/genLoginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).GenLoginInfo(ctx, req.(*LoginInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_BlockConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).BlockConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/blockConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).BlockConversation(ctx, req.(*BlockConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalAccountsService_GetConversationBlockStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationBlockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalAccountsServiceServer).GetConversationBlockStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InternalAccountsService/getConversationBlockStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalAccountsServiceServer).GetConversationBlockStatus(ctx, req.(*GetConversationBlockReq))
	}
	return interceptor(ctx, in, info, handler)
}

// InternalAccountsService_ServiceDesc is the grpc.ServiceDesc for InternalAccountsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InternalAccountsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "InternalAccountsService",
	HandlerType: (*InternalAccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getInfo",
			Handler:    _InternalAccountsService_GetInfo_Handler,
		},
		{
			MethodName: "getInfoByEmail",
			Handler:    _InternalAccountsService_GetInfoByEmail_Handler,
		},
		{
			MethodName: "disable",
			Handler:    _InternalAccountsService_Disable_Handler,
		},
		{
			MethodName: "enable",
			Handler:    _InternalAccountsService_Enable_Handler,
		},
		{
			MethodName: "getAll",
			Handler:    _InternalAccountsService_GetAll_Handler,
		},
		{
			MethodName: "edit",
			Handler:    _InternalAccountsService_Edit_Handler,
		},
		{
			MethodName: "renew",
			Handler:    _InternalAccountsService_Renew_Handler,
		},
		{
			MethodName: "createAccount",
			Handler:    _InternalAccountsService_CreateAccount_Handler,
		},
		{
			MethodName: "queryAccount",
			Handler:    _InternalAccountsService_QueryAccount_Handler,
		},
		{
			MethodName: "upload",
			Handler:    _InternalAccountsService_Upload_Handler,
		},
		{
			MethodName: "uploadAvatar",
			Handler:    _InternalAccountsService_UploadAvatar_Handler,
		},
		{
			MethodName: "kickOffDevice",
			Handler:    _InternalAccountsService_KickOffDevice_Handler,
		},
		{
			MethodName: "getUserTeams",
			Handler:    _InternalAccountsService_GetUserTeams_Handler,
		},
		{
			MethodName: "syncAccountBuInfo",
			Handler:    _InternalAccountsService_SyncAccountBuInfo_Handler,
		},
		{
			MethodName: "inactive",
			Handler:    _InternalAccountsService_Inactive_Handler,
		},
		{
			MethodName: "downloadAvatar",
			Handler:    _InternalAccountsService_DownloadAvatar_Handler,
		},
		{
			MethodName: "genLoginInfo",
			Handler:    _InternalAccountsService_GenLoginInfo_Handler,
		},
		{
			MethodName: "blockConversation",
			Handler:    _InternalAccountsService_BlockConversation_Handler,
		},
		{
			MethodName: "getConversationBlockStatus",
			Handler:    _InternalAccountsService_GetConversationBlockStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "InternalAccountsService.proto",
}
