package logic

import (
	"context"

	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisableSearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDisableSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisableSearchLogic {
	return &DisableSearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DisableSearchLogic) DisableSearch(in *pb.DisableSearchRequest) (*pb.DisableSearchResponse, error) {
	base := &pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	var hKey = logic.DirectoryHashKey
	var db = l.svcCtx.DbEngine
	// 更新redis
	redisCmd := l.svcCtx.Redis
	var profile = models.UserProfile{UID: in.Uid}
	db.Model(profile).Where(profile).First(&profile)
	if profile.EmailHash != "" {
		redisCmd.HDel(l.ctx, hKey, profile.EmailHash)
	}
	if profile.PhoneHash != "" {
		redisCmd.HDel(l.ctx, hKey, profile.PhoneHash)
	}

	l.Infow("DisableSearch", logx.Field("uid", in.Uid))

	return &pb.DisableSearchResponse{Base: base}, nil
}
