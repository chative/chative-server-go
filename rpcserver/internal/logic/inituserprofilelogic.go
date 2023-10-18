package logic

import (
	"context"
	"errors"

	"chative-server-go/models"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInitUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitUserProfileLogic {
	return &InitUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InitUserProfileLogic) InitUserProfile(in *pb.InitUserProfileRequest) (*pb.InitUserProfileResponse, error) {
	respBase := pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	resp := &pb.InitUserProfileResponse{Base: &respBase}
	db := l.svcCtx.DbEngine
	var profile = models.UserProfile{UID: in.Uid}
	err := db.Model(profile).Where(profile).First(&profile).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorw("InitUserProfile failed,first record failed", logx.Field("uid", in.Uid), logx.Field("err", err))
		resp.Base.Status = 1
		resp.Base.Reason = err.Error()
		return resp, nil
	}
	profile.EmailHash = in.EmailHash
	profile.PhoneHash = in.PhoneHash
	err = db.Save(&profile).Error
	if err != nil {
		l.Errorw("InitUserProfile failed,save failed", logx.Field("uid", in.Uid), logx.Field("err", err))
		resp.Base.Status = 1
		resp.Base.Reason = err.Error()
		return resp, nil
	}
	l.Infow("InitUserProfile success", logx.Field("autoid", profile.ID), logx.Field("uid", in.Uid), logx.Field("EmailHash", in.EmailHash), logx.Field("PhoneHash", in.PhoneHash))
	return resp, nil
}
