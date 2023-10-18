package logic

import (
	"context"
	"errors"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/models"
	"chative-server-go/response"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile(req *types.GetProfileReq) (resp *types.ProfileResp, errInfo *response.ErrInfo) {
	var profile = models.UserProfile{UID: req.UID}
	resp = &types.ProfileResp{}
	db := l.svcCtx.DbEngine
	err := db.Where(profile).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		l.Errorw("GetProfile", logx.Field("req", req), logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	resp.SearchByEmail = profile.SearchByEmail
	resp.SearchByPhone = profile.SearchByPhone
	// 查询是否有passkeys
	var webauthnUser = models.WebauthnUser{ChatUID: req.UID}
	err = db.Where(webauthnUser).First(&webauthnUser).Error
	if err == nil && webauthnUser.CredentialCnt > 0 {
		resp.PasskeysSwitch = 1
	} else {
		resp.PasskeysSwitch = 0
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			l.Errorw("GetProfile PasskeysSwitch error", logx.Field("err", err))
		}
	}
	return
}
