package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httputil"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/mainrpc"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/gofrs/uuid"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebauthnLoginFinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebauthnLoginFinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebauthnLoginFinLogic {
	return &WebauthnLoginFinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebauthnLoginFinLogic) WebauthnLoginFin(req *types.WebauthnLoginFinReq) (resp *types.WebauthnLoginFinRes, errInfo *response.ErrInfo) {
	cfg := l.svcCtx.Config
	reqJson, err := json.Marshal(req)
	if err != nil {
		l.Errorw("WebauthnLoginFin marshal failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	reqHttp, err := http.NewRequest(http.MethodPost, "http://"+cfg.Webauthn.Host+"/webauthn/login/finalize", bytes.NewReader(reqJson))
	if err != nil {
		l.Errorw("WebauthnLoginFin new request failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	reqHttp.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(reqHttp)
	if err != nil {
		l.Errorw("WebauthnLoginFin do request failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		resBody, _ := httputil.DumpResponse(res, true)
		l.Errorw("WebauthnLoginFin do request failed",
			logx.Field("status", res.StatusCode), logx.Field("req", string(reqJson)),
			logx.Field("res", string(resBody)))
		errInfo = &response.ErrInfo{
			ErrCode: 12003,
			Reason:  "Unauthorized",
		}
		return
	}
	var httpRes = struct {
		UserID string `json:"user_id"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&httpRes)
	if err != nil {
		l.Errorw("WebauthnLoginFin decode response failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	// 查对应的uid
	var webauthnUser = models.WebauthnUser{UserID: uuid.FromStringOrNil(httpRes.UserID)}
	err = l.svcCtx.DbEngine.First(&webauthnUser).Error
	if err != nil {
		l.Errorw("WebauthnLoginFin find webauthn user failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	loginInfo, err := mainrpc.GenLoginInfo(webauthnUser.ChatUID, req.UA, req.SupportTransfer != 0)
	if err != nil {
		l.Errorw("WebauthnLoginFin gen login info failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	resp = new(types.WebauthnLoginFinRes)
	err = json.Unmarshal(loginInfo, resp)
	l.Infow("WebauthnLoginFin,after GenLoginInfo", logx.Field("loginfo", loginInfo), logx.Field("resp", resp))
	if err != nil {
		l.Errorw("WebauthnLoginFin unmarshal login info failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
	}
	return
}
