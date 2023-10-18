// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"chative-server-go/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPut,
				Path:    "/v1/friend",
				Handler: addFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v1/utils/location",
				Handler: getLocationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v3/friend/ask",
				Handler: askNewFriendHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/v3/friend/ask/:id/agree",
				Handler: friendAgreeHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/v3/friend/:uid",
				Handler: deleteFriendHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v1/conversationconfig/share",
				Handler: fetchConversationConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/v1/conversationconfig/share/:id",
				Handler: updateConversationConfigHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v3/directory/searchsalt",
				Handler: fetchSaltHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v3/directory/search",
				Handler: searchContactsHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/v3/directory/profile",
				Handler: setProfileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v3/directory/profile",
				Handler: getProfileHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v3/keys/identity/bulk",
				Handler: QueryAccIdKeysHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v3/accounts/inviteCode",
				Handler: GenInviteCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/v3/accounts/querybyInviteCode",
				Handler: QueryByInviteCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v3/accounts/exists",
				Handler: ExistsUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v3/accounts/report",
				Handler: ReportAccHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v3/web/userinfo",
				Handler: UserInfoByInviteCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v3/web/linkInvitationInfo",
				Handler: LinkInvitationInfoHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/v3/upgrade/android/check",
				Handler: checkAndroidUpdateHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v3/webauthn/user",
				Handler: webauthnExistsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v3/webauthn/registration/initialize",
				Handler: webauthnRegInitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v3/webauthn/registration/finalize",
				Handler: webauthnRegFinHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v3/webauthn/login/initialize",
				Handler: webauthnLoginInitHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/v3/webauthn/login/finalize",
				Handler: webauthnLoginFinHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/v3/conversationconfig/reminder",
				Handler: createReminderHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/v3/conversationconfig/reminder/:id",
				Handler: updateReminderHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/v3/conversationconfig/reminder/:ids",
				Handler: deleteReminderHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/v3/conversationconfig/reminder/:conversation",
				Handler: getReminderHandler(serverCtx),
			},
		},
	)
}
