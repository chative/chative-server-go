package logic

import (
	"context"
	"strings"

	"chative-server-go/rediscluster"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"
	"chative-server-go/utils/crypto"
	"chative-server-go/utils/secretsmanager"

	"github.com/go-redis/redis_rate/v9"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckEmailVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// 0. 去掉空格，转换为小写
func ClearEmail(id string) string {
	return strings.ToLower(strings.TrimSpace(id))
}

func vCodeEmailKey(hash string) string {
	return "vCode:email:login:" + hash
}

func NewCheckEmailVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEmailVerificationCodeLogic {
	return &CheckEmailVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckEmailVerificationCodeLogic) CheckEmailVerificationCode(in *pb.CheckEmailVcodeRequest) (*pb.CheckEmailVcodeResponse, error) {
	base := &pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	var hash = in.EmailHash
	var err error

	if hash == "" {
		// 0. 去掉空格，转换为小写
		email := ClearEmail(in.Email)
		// 1. 生成hash
		sm := secretsmanager.GetSM()
		hash, err = crypto.HashID(email, sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			base.Status = 99
			base.Reason = "Internal error, please try again later.(HashFailed)"
			l.Errorw("CheckEmailVerificationCode", logx.Field("err", err))
			return &pb.CheckEmailVcodeResponse{Base: base}, nil
		}
	}
	limiter := redis_rate.NewLimiter(l.svcCtx.Redis)

	res, err := limiter.Allow(l.ctx, strings.Join([]string{"limit:verifycode", hash}, ":"), rediscluster.VerificationCodeLimit)
	if err != nil {
		l.Errorw("CheckEmailVerificationCode Rate Limit Failed", logx.Field("err", err))
		base.Status = 99
		base.Reason = "Internal Server Error"
		return &pb.CheckEmailVcodeResponse{Base: base}, nil
	}
	if res.Allowed < 1 {
		l.Errorw("CheckEmailVerificationCode Rate Limit Exceeded", logx.Field("err", err))
		base.Status = 27
		base.Reason = "Please resend and try again after 1 minute."
		return &pb.CheckEmailVcodeResponse{Base: base}, nil
	}
	// 2. redis中取出验证码
	redisCmd := l.svcCtx.Redis
	vCode, err := redisCmd.Get(l.ctx, vCodeEmailKey(hash)).Result()
	if err != nil {
		base.Status = 99
		base.Reason = "Internal error, please try again later.(FindVCodeFailed)"
		l.Errorw("CheckEmailVerificationCode", logx.Field("hash", hash), logx.Field("err", err))
		return &pb.CheckEmailVcodeResponse{Base: base}, nil
	}
	// 3. 比较
	if vCode != in.Vcode {
		base.Status = 1
		base.Reason = "Verification code error."
		l.Errorw("CheckEmailVerificationCode", logx.Field("hash", hash), logx.Field("err", err), logx.Field("vCode", vCode))
		return &pb.CheckEmailVcodeResponse{Base: base}, nil
	} else {
		// 4. 删除验证码
		_, err = redisCmd.Del(l.ctx, vCodeEmailKey(hash)).Result()
		if err != nil {
			l.Errorw("CheckEmailVerificationCode", logx.Field("hash", hash), logx.Field("err", err))
			return &pb.CheckEmailVcodeResponse{Base: base}, nil
		}
	}

	return &pb.CheckEmailVcodeResponse{Base: base}, nil
}
