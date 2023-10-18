package logic

import (
	"context"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	"chative-server-go/response"
	"chative-server-go/utils/crypto"
	"chative-server-go/utils/secretsmanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchContactsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchContactsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchContactsLogic {
	return &SearchContactsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchContactsLogic) SearchContacts(req *types.DirectorySearchReq) (resp *types.DirectorySearchRes, errInfo *response.ErrInfo) {
	redisCmd := l.svcCtx.Redis
	serverHashs := make([]string, 0, len(req.Hashs))
	sm := secretsmanager.GetSM()
	for _, hash := range req.Hashs {
		serverHash, err := crypto.ServerHash(hash, sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("ServerHash error", logx.Field("hash", hash), logx.Field("err", err))
			continue
		}
		serverHashs = append(serverHashs, serverHash)
	}
	arr, err := redisCmd.HMGet(l.ctx, logic.DirectoryHashKey, serverHashs...).Result()
	if err != nil {
		l.Errorw("SearchContacts", logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		return
	}
	var results = make([]DirectoryResult, 0, len(arr))
	for i, v := range req.Hashs {
		if arr[i] == nil {
			continue
		}
		l.Info("SearchContacts", logx.Field("hash", v), logx.Field("uid", arr[i]))
		if uid, ok := arr[i].(string); ok {
			results = append(results, DirectoryResult{Hash: v, UID: uid})
		}
	}
	if len(results) == 0 {
		l.Infow("SearchContacts", logx.Field("req", req), logx.Field("results", results), logx.Field("serverHashs", serverHashs), logx.Field("arr", arr))
		return
	}

	if len(results) == 1 {
		logic.CheckFindPath(logic.NewContext(l.ctx, l.Logger, l.svcCtx.DbEngine, l.svcCtx.Redis),
			req.UID, []string{results[0].UID}, "search")
	}

	resp = &types.DirectorySearchRes{Results: results}
	return
}

type DirectoryResult struct {
	Hash string `json:"hash"`
	UID  string `json:"uid"`
}
