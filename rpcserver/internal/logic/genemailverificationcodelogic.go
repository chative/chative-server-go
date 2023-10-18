package logic

import (
	"context"
	"time"

	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"
	"chative-server-go/utils/crypto"
	"chative-server-go/utils/secretsmanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenEmailVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenEmailVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenEmailVerificationCodeLogic {
	return &GenEmailVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenEmailVerificationCodeLogic) GenEmailVerificationCode(in *pb.GenEmailVcodeRequest) (*pb.GenEmailVcodeResponse, error) {
	base := &pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	// 0. 去掉空格，转换为小写
	email := ClearEmail(in.Email)
	// 1. 生成hash
	sm := secretsmanager.GetSM()
	hash, err := crypto.HashID(email, sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
	if err != nil {
		base.Status = 99
		base.Reason = "Internal error, please try again later.(HashFailed)"
		l.Errorw("CheckEmailVerificationCode", logx.Field("err", err), logx.Field("email", email))
		return &pb.GenEmailVcodeResponse{Base: base}, nil
	}
	// 2. 生成随机数验证码
	vCode := crypto.GenRandomNumber(6)
	// 3. 写到redis
	redisCmd := l.svcCtx.Redis
	res, err := redisCmd.SetEX(l.ctx, vCodeEmailKey(hash), vCode, time.Hour).Result()
	l.Infow("GenEmailVerificationCode", logx.Field("hash", hash), logx.Field("res", res), logx.Field("err", err))
	if err != nil {
		base.Status = 99
		base.Reason = "Internal error, please try again later.(SaveVCodeFailed)"
		l.Errorw("CheckEmailVerificationCode", logx.Field("err", err), logx.Field("email", email))
		return &pb.GenEmailVcodeResponse{Base: base}, nil
	}

	return &pb.GenEmailVcodeResponse{Base: base, Vcode: vCode}, nil
}
