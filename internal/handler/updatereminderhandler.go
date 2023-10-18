package handler

import (
	"net/http"

	"chative-server-go/internal/logic"
	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/response" //

	"github.com/zeromicro/go-zero/rest/httpx"
)

func updateReminderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateReminderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUpdateReminderLogic(r.Context(), svcCtx)
		resp, err := l.UpdateReminder(&req)
		response.Response(w, resp, err) //②

	}
}
