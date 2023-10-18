package logic

import (
	"context"

	"chative-server-go/models"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"
	"chative-server-go/utils/crypto"
	"chative-server-go/utils/secretsmanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserProfileLogic) GetUserProfile(in *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	base := &pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	var err error
	sm := secretsmanager.GetSM()
	var db = l.svcCtx.DbEngine
	var profile models.UserProfile
	conditon := models.UserProfile{UID: in.Uid}
	if in.Email != "" {
		conditon.EmailHash, err = crypto.HashID(ClearEmail(in.Email), sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("GetUserProfile Failed,HashID(Failed)", logx.Field("err", err))
		}
	} else if in.EmailHash != "" {
		conditon.EmailHash = in.EmailHash
	}

	if in.Phone != "" {
		conditon.PhoneHash, err = crypto.HashID(in.Phone, sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("GetUserProfile Failed,HashID(Failed)", logx.Field("err", err))
		}
	} else if in.PhoneHash != "" {
		conditon.PhoneHash = in.PhoneHash
	}

	err = db.Where(conditon).First(&profile).Error
	resp := &pb.GetUserProfileResponse{Base: base}
	if err != nil {
		base.Status = 404
		base.Reason = err.Error()
		l.Errorw("GetUserProfile Failed", logx.Field("err", err), logx.Field("conditon", conditon))
	} else {
		resp.EmailHash = profile.EmailHash
		resp.PhoneHash = profile.PhoneHash
		resp.Uid = profile.UID
	}
	return resp, nil
}
