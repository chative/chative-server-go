package logic

import (
	"context"

	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"
	"chative-server-go/utils/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSMSLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendSMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSMSLogic {
	return &SendSMSLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendSMSLogic) SendSMS(in *pb.SendSMSRequest) (*pb.SendSMSResponse, error) {
	err := sms.SendCode(in.Phone)
	base := &pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	if err == nil {
		return &pb.SendSMSResponse{Base: base}, nil
	}
	l.Errorw("SendSMS", logx.Field("err", err), logx.Field("phone", in.Phone))
	if err == sms.ErrNotSupport {
		base.Status = 10006
		base.Reason = "Phone numbers in this area are not supported."
		return &pb.SendSMSResponse{Base: base}, nil
	}
	base.Status = 10007
	base.Reason = "Send SMS failed"
	return &pb.SendSMSResponse{Base: base}, nil
}
