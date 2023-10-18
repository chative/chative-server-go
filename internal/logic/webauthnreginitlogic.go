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
	"chative-server-go/utils/crypto"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebauthnRegInitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebauthnRegInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebauthnRegInitLogic {
	return &WebauthnRegInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebauthnRegInitLogic) WebauthnRegInit(req *types.WebauthnRegInitReq) (resp *types.WebauthnRegInitRes, errInfo *response.ErrInfo) {
	// get user(存在就删除，不存在就创建)
	db := l.svcCtx.DbEngine
	cfg := l.svcCtx.Config
	var webauthnUser = models.WebauthnUser{ChatUID: req.ChatUID}
	err := db.First(&webauthnUser, webauthnUser).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		mockEmail := req.ChatUID + crypto.GenRandomString(8) + "@chat.internal"
		userID, password, err2 := logic.CreateWebauthnUser(cfg.Webauthn.Host, mockEmail)
		if err2 != nil {
			l.Errorw("WebauthnRegInit create user failed", logx.Field("err", err2))
			errInfo = &response.ErrInfo{
				ErrCode: 99,
				Reason:  "internal error",
			}
			return
		}
		webauthnUser.UserID = uuid.FromStringOrNil(userID) //userID
		webauthnUser.Password = password
		err = db.Create(&webauthnUser).Error
	}
	if err != nil {
		l.Error("WebauthnRegInit get user failed", logx.Field("err", err))
	}
	// 调用webauthn的注册初始化
	respBody, err := logic.WebauthnRegInit(cfg.Webauthn.Host, webauthnUser.UserID.String(),
		webauthnUser.Password)
	if err != nil {
		l.Errorw("WebauthnRegInit failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	resp = &types.WebauthnRegInitRes{}
	json.Unmarshal(respBody, resp)
	return
}
