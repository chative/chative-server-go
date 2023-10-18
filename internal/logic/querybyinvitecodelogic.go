package logic

import (
	"context"
	"errors"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/response"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryByInviteCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryByInviteCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByInviteCodeLogic {
	return &QueryByInviteCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryByInviteCodeLogic) QueryByInviteCode(req *types.QueryByInviteCodeReq) (resp *types.QueryByInviteCodeRes, errInfo *response.ErrInfo) {
	db := l.svcCtx.DbEngine
	var accountInternal = models.InternalAccount{InviteCode: &req.InviteCode}
	*accountInternal.InviteCode = (*accountInternal.InviteCode)[:8]
	err := db.Where(accountInternal).First(&accountInternal).Error

	if err == nil {
		resp = &types.QueryByInviteCodeRes{
			UID: accountInternal.Number,
		}
		logic.CheckFindPath(logic.NewContext(l.ctx, l.Logger, l.svcCtx.DbEngine, l.svcCtx.Redis),
			req.UID, []string{resp.UID}, "link")
		return
	}
	l.Errorw("QueryByInviteCode failed",
		logx.Field("inviteCode", req.InviteCode), logx.Field("err", err))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errInfo = &response.ErrInfo{
			ErrCode: 10107,
			Reason:  "Expired Invite Code",
		}
		return
	}

	errInfo = &response.ErrInfo{
		ErrCode: 99,
		Reason:  "Expired Invite Code",
	}
	return
}
