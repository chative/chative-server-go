package logic

import (
	"context"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/response"
	"chative-server-go/utils/secretsmanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type FetchSaltLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFetchSaltLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FetchSaltLogic {
	return &FetchSaltLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FetchSaltLogic) FetchSalt(req *types.FetchSaltReq) (resp *types.FetchSaltResp, errInfo *response.ErrInfo) {
	resp = &types.FetchSaltResp{Salt: secretsmanager.GetSM().GetDirectoryClientSalt()}
	return
}
