package logic

import (
	"context"

	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserHashLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserHashLogic {
	return &DelUserHashLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserHashLogic) DelUserHash(in *pb.DelUserHashRequest) (*pb.DelUserHashResponse, error) {
	base := &pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	var hKey = logic.DirectoryHashKey
	var db = l.svcCtx.DbEngine
	// 更新redis
	redisCmd := l.svcCtx.Redis
	var profile = models.UserProfile{UID: in.Uid}
	db.Model(profile).Where(profile).First(&profile)
	updates := map[string]interface{}{}
	if in.DelEmail {
		updates["email_hash"] = ""
		if profile.EmailHash != "" {
			redisCmd.HDel(l.ctx, hKey, profile.EmailHash)
		}
	}
	if in.DelPhone {
		updates["phone_hash"] = ""
		if profile.PhoneHash != "" {
			redisCmd.HDel(l.ctx, hKey, profile.PhoneHash)
		}
	}
	if len(updates) == 0 {
		l.Infow("DelUserHash Failed,Nothing to update", logx.Field("in", in))
		return &pb.DelUserHashResponse{Base: base}, nil
	}

	res := db.Model(models.UserProfile{}).Where(models.UserProfile{UID: in.Uid}).Updates(updates)
	l.Infow("DelUserHash", logx.Field("uid", in.Uid), logx.Field("updates", updates), logx.Field("RowsAffected", res.RowsAffected))
	err := res.Error
	if err != nil {
		l.Errorw("DelUserHash failed", logx.Field("uid", in.Uid), logx.Field("err", err))
	}

	return &pb.DelUserHashResponse{Base: base}, nil
}
