package handler

import (
	"net/http"

	"chative-server-go/internal/logic"
	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/response" //

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LinkInvitationInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	logic.InitOSSclient(&svcCtx.Config)
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupInviteInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLinkInvitationInfoLogic(r.Context(), svcCtx)
		resp, err := l.LinkInvitationInfo(&req)
		response.Response(w, resp, err) //②

	}
}
