package logic

import (
	"context"
	"errors"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/models"
	"chative-server-go/response"
	"chative-server-go/utils/crypto"
	"chative-server-go/utils/secretsmanager"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExistsUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExistsUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistsUserLogic {
	return &ExistsUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExistsUserLogic) ExistsUser(req *types.CheckAccExistsReq) (resp *types.CheckAccExistsResp, errInfo *response.ErrInfo) {
	// 1. 计算uid ; 先算email，再算phone
	if req.Email == "" && req.Phone == "" {
		errInfo = &response.ErrInfo{
			ErrCode: 1,
			Reason:  "Invalid parameter,email and phone is empty",
		}
		return
	}
	var err error
	var profile models.UserProfile
	sm := secretsmanager.GetSM()
	if req.Email != "" {
		profile.EmailHash, err = crypto.HashID(ClearEmail(req.Email), sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("WebauthnExists hash email failed", logx.Field("err", err))
			errInfo = &response.ErrInfo{
				ErrCode: 99,
				Reason:  "internal error",
			}
			return
		}
	} else if req.Phone != "" {
		profile.PhoneHash, err = crypto.HashID(req.Phone, sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("WebauthnExists hash phone failed", logx.Field("err", err))
			errInfo = &response.ErrInfo{
				ErrCode: 99,
				Reason:  "internal error",
			}
			return
		}
	}

	// 2. 查询数据库
	var db = l.svcCtx.DbEngine
	err = db.First(&profile, profile).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		l.Logger.Infow("WebauthnExists user not found", logx.Field("profile", profile))
		resp = &types.CheckAccExistsResp{
			Exists: false,
		}
		return
	}
	if err != nil {
		l.Errorw("WebauthnExists profile db error", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	var webauthn models.WebauthnUser
	webauthn.ChatUID = profile.UID
	err = db.First(&webauthn, webauthn).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		l.Logger.Infow("WebauthnExists, webauthn not found", logx.Field("profile", profile))
		resp = &types.CheckAccExistsResp{
			Exists: true, HasWebauthn: false,
		}
		return
	}
	if err != nil {
		l.Errorw("WebauthnExists webauthn db error", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	if webauthn.CredentialCnt == 0 {
		l.Logger.Infow("WebauthnExists, webauthn credentials is 0", logx.Field("profile", profile))
		resp = &types.CheckAccExistsResp{
			Exists: true, HasWebauthn: false,
		}
		return
	}
	resp = &types.CheckAccExistsResp{
		Exists: true, HasWebauthn: true,
		WebauthnUserID: webauthn.UserID.String(),
	}

	return
}
