package logic

import (
	"context"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/models"
	"chative-server-go/response"
	"chative-server-go/utils/crypto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenInviteCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenInviteCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenInviteCodeLogic {
	return &GenInviteCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenInviteCodeLogic) GenInviteCode(req *types.GenInviteCodeReq) (resp *types.GenInviteCodeRes, errInfo *response.ErrInfo) {
	//
	db := l.svcCtx.DbEngine
	var internalAcc = models.InternalAccount{Number: req.UID}
	err := db.Where(internalAcc).First(&internalAcc).Error
	if err != nil {
		l.Errorw("GenInviteCode failed,query account failed",
			logx.Field("uid", req.UID), logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "GenInviteCode failed",
		}
		return
	}
	if internalAcc.InviteCode == nil || len(*internalAcc.InviteCode) == 0 {
		req.Regenerate = 1
	}
	if req.Regenerate == 0 {
		resp = &types.GenInviteCodeRes{
			InviteCode: *internalAcc.InviteCode,
		}
		if req.Short == 0 {
			resp.InviteCode += crypto.GenRandomString(24)
		}
		return
	}
	for i := 0; i < 3; i++ {
		// 生成邀请码
		inviteCode := crypto.GenRandomString(8)
		// 保存到数据库
		result := db.Model(internalAcc).Where(models.InternalAccount{Number: internalAcc.Number}).
			Update("invite_code", inviteCode)
		if result.Error != nil {
			l.Errorw("GenInviteCode failed,update account failed",
				logx.Field("uid", req.UID), logx.Field("err", result.Error))
			continue
		}
		if result.RowsAffected > 0 {
			resp = &types.GenInviteCodeRes{
				InviteCode: inviteCode,
			}
			if req.Short == 0 {
				resp.InviteCode += crypto.GenRandomString(24)
			}
			return
		}
	}
	errInfo = &response.ErrInfo{
		ErrCode: 99,
		Reason:  "GenInviteCode failed",
	}
	return
}
