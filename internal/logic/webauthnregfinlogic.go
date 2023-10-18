package logic

import (
	"context"
	"encoding/json"
	"errors"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/response"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebauthnRegFinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebauthnRegFinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebauthnRegFinLogic {
	return &WebauthnRegFinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebauthnRegFinLogic) WebauthnRegFin(req *types.WebauthnRegFinReq) (resp *types.WebauthnRegFinRes, errInfo *response.ErrInfo) {

	db := l.svcCtx.DbEngine
	cfg := l.svcCtx.Config
	var webauthnUser = models.WebauthnUser{ChatUID: req.ChatUID}
	err := db.First(&webauthnUser, webauthnUser).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		errInfo = &response.ErrInfo{
			ErrCode: 12001,
			Reason:  "not found",
		}
		return
	}
	// 调用webauthn的注册结束
	regFinReq, err := json.Marshal(req)
	if err != nil {
		l.Errorw("WebauthnRegFin marshal failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	httpCode, err := logic.WebauthnRegFin(l, cfg.Webauthn.Host, webauthnUser.UserID.String(),
		webauthnUser.Password, regFinReq)
	if httpCode >= 300 {
		l.Errorw("WebauthnRegFin failed", logx.Field("httpCode", httpCode))
		errInfo = &response.ErrInfo{
			ErrCode: 12002,
			Reason:  "webauthn register failed",
		}
		return
	}
	if err != nil {
		l.Errorw("WebauthnRegFin failed", logx.Field("err", err), logx.Field("httpCode", httpCode))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	// Credential的个数置为1
	err = db.Model(&webauthnUser).Where(models.WebauthnUser{ChatUID: req.ChatUID}).Update("credential_cnt", 1).Error
	if err != nil {
		l.Errorw("WebauthnRegFin update credential_cnt failed",
			logx.Field("err", err), logx.Field("uid", req.ChatUID))
	}

	return
}

func DelWebauthnCredentials(db *gorm.DB, host, chatUID string) (err error) {
	var webauthnUser = models.WebauthnUser{ChatUID: chatUID}
	err = db.First(&webauthnUser, webauthnUser).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	err = logic.WebauthnDeleteCredentials(host, webauthnUser.UserID.String(),
		webauthnUser.Password)
	if err != nil {
		return
	}
	// Credential的个数置为0
	err = db.Model(&webauthnUser).Where(models.WebauthnUser{ChatUID: chatUID}).
		Update("credential_cnt", 0).Error
	return
}
