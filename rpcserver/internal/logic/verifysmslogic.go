package logic

import (
	"context"
	"strings"

	"chative-server-go/rediscluster"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"
	"chative-server-go/utils/crypto"
	"chative-server-go/utils/secretsmanager"
	"chative-server-go/utils/sms"

	"github.com/go-redis/redis_rate/v9"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifySMSLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifySMSLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifySMSLogic {
	return &VerifySMSLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifySMSLogic) VerifySMS(in *pb.VerifySMSRequest) (*pb.VerifySMSResponse, error) {
	base := &pb.BaseResponse{Ver: 1, Status: 0, Reason: "OK"}
	limiter := redis_rate.NewLimiter(l.svcCtx.Redis)
	sm := secretsmanager.GetSM()
	hash, err := crypto.HashID(in.Phone, sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
	if err != nil {
		l.Errorw("VerifySMS HashID Failed", logx.Field("err", err))
		base.Status = 99
		base.Reason = "Internal Server Error"
		return &pb.VerifySMSResponse{Base: base}, nil
	}
	res, err := limiter.Allow(l.ctx, strings.Join([]string{"limit:verifycode", hash}, ":"), rediscluster.VerificationCodeLimit)
	if err != nil {
		l.Errorw("VerifySMS Rate Limit Failed", logx.Field("err", err))
		base.Status = 99
		base.Reason = "Internal Server Error"
		return &pb.VerifySMSResponse{Base: base}, nil
	}
	if res.Allowed < 1 {
		l.Errorw("VerifySMS Rate Limit Exceeded", logx.Field("err", err))
		base.Status = 27
		base.Reason = "Please resend and try again after 1 minute."
		return &pb.VerifySMSResponse{Base: base}, nil
	}
	err = sms.VerifyCode(in.Phone, in.Code)
	if err == nil {
		return &pb.VerifySMSResponse{Base: base}, nil
	}
	l.Errorw("VerifySMS", logx.Field("err", err), logx.Field("phoneHash", hash), logx.Field("code", in.Code))
	base.Status = 10008
	base.Reason = "Verification failed"
	return &pb.VerifySMSResponse{Base: base}, nil
}
