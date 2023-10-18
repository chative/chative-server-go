package logic

import (
	"context"
	"strings"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchConversationConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchConversationConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchConversationConfigLogic {
	return &FetchConversationConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchConversationConfigLogic) FetchConversationConfig(req *types.FetchConversationRequest) (resp *types.FetchConversationResponse, errInfo *response.ErrInfo) {
	if len(req.Conversations) > 50 { // 一次最多50个
		errInfo = &response.ErrInfo{
			ErrCode: 1,
			Reason:  "Invalid parameter",
		}
		l.Errorw("FetchConversationConfig invalid parameter", logx.Field("req", req))
		return
	}
	for _, v := range req.Conversations {
		if arr := strings.Split(v, ":"); len(arr) != 2 || arr[0] != req.UID && arr[1] != req.UID {
			errInfo = &response.ErrInfo{
				ErrCode: 1,
				Reason:  "Invalid parameter",
			}
			l.Errorw("FetchConversationConfig invalid conversation", logx.Field("conversation", v))
			return
		}
	}
	db := l.svcCtx.DbEngine
	var configs []models.ShareConversationCnf
	err := db.Model(models.ShareConversationCnf{}).Where("conversation in ?", req.Conversations).
		Find(&configs).Error
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Internal error",
		}
		l.Errorw("FetchConversationConfig db error", logx.Field("err", err))
		return
	}
	resp = &types.FetchConversationResponse{Conversations: make([]types.Conversation, 0, len(configs))}
	for _, v := range configs {
		resp.Conversations = append(resp.Conversations, types.Conversation{
			Ver:           v.Version,
			Conversation:  v.Conversation,
			MessageExpiry: v.MessageExpiry,
		})
	}
	return
}
