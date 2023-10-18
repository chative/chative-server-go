package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebauthnLoginInitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebauthnLoginInitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebauthnLoginInitLogic {
	return &WebauthnLoginInitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebauthnLoginInitLogic) WebauthnLoginInit(req *types.WebauthnLoginInitReq) (resp *types.WebauthnLoginInitRes, errInfo *response.ErrInfo) {
	cfg := l.svcCtx.Config.Webauthn
	// 调用webauthn的登录初始化
	var reqJson = struct {
		UserID string `json:"user_id"`
	}{req.UserID}
	reqBody, err := json.Marshal(reqJson)
	if err != nil {
		l.Errorw("WebauthnLoginInit marshal failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	res, err := http.Post("http://"+cfg.Host+"/webauthn/login/initialize", "application/json", bytes.NewReader(reqBody))
	// reqInit, err := http.NewRequest(http.MethodPost, "http://"+cfg.Host+"/webauthn/login/initialize", bytes.NewReader(reqBody))
	// if err != nil {
	// 	l.Errorw("WebauthnLoginInit new request failed", logx.Field("err", err))
	// 	errInfo = &response.ErrInfo{
	// 		ErrCode: 99,
	// 		Reason:  "internal error",
	// 	}
	// 	return
	// }
	// reqInit.Header.Set("Content-Type", "application/json")
	// res, err := http.DefaultClient.Do(reqInit)
	if err != nil {
		l.Errorw("WebauthnLoginInit do request failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		// reqDump, _ := httputil.DumpRequest(reqInit, true)
		l.Errorw("WebauthnLoginInit do request failed",
			logx.Field("status", res.StatusCode), logx.Field("req", string(reqBody)))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		l.Errorw("WebauthnLoginInit decode response failed", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	return
}
