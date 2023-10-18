package logic

import (
	"context"
	"time"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/rediscluster"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddFriendLogic) AddFriend(req *types.AddFriendReq) (resp *types.AddFriendRes, errInfo *response.ErrInfo) {
	var invite models.Invitation
	tNow := time.Now()
	db := l.svcCtx.DbEngine
	err := db.Where("code = ?", req.InviteCode).First(&invite).Error
	// 1. invite Code是自己的直接返回成功。(client跳到自己的note会话)
	if err == nil && invite.Inviter == req.UID {
		return &types.AddFriendRes{Inviter: invite.Inviter}, nil
	}

	// 检查inviteCode是否不存在、过期、被邀请人invalid、被使用过，检查失败就返回错误
	if err != nil || invite.Timestamp+86400*1000*7 < tNow.UnixMilli() || invite.RegisterTime > 0 {
		errInfo = &response.ErrInfo{
			ErrCode: 19001,
			Reason:  "Invalid invite code.",
		}
		l.Infow("Invalid invite code.", logx.Field("inviteCode", req.InviteCode),
			logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}

	logicCtx := logic.NewContext(l.ctx, l.Logger, db, rediscluster.GetRedis())
	// 被block了
	blockedInvitee, err := logic.InBlockList(logicCtx, invite.Inviter, req.UID)
	if blockedInvitee {
		return &types.AddFriendRes{Inviter: invite.Inviter}, nil
	}
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Unkown Server Error.",
		}
		l.Infow("InBlockList Exception", logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}
	// 扫了黑名单里面的人，提示unblock
	blockedInviter, err := logic.InBlockList(logicCtx, req.UID, invite.Inviter)
	if blockedInviter {
		resp = &types.AddFriendRes{Inviter: invite.Inviter}
		errInfo = &response.ErrInfo{
			ErrCode: 19010,
			Reason:  "You have blocked the inviter.",
		}
		return
	}
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Unkown Server Error.",
		}
		l.Infow("InBlockList Exception", logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}

	// 占用inviteCode
	res := db.Model(models.Invitation{}).Where(models.Invitation{Code: req.InviteCode}).Where("register_time=0").
		UpdateColumns(models.Invitation{RegisterTime: tNow.UnixMilli(), Account: req.UID})
	if res.RowsAffected == 0 {
		errInfo = &response.ErrInfo{
			ErrCode: 19001,
			Reason:  "This invite code has been used.",
		}
		l.Infow("Invalid invite code.", logx.Field("inviteCode", req.InviteCode),
			logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}

	logic.AddFriend(logicCtx,
		&logic.AddReq{Inviter: invite.Inviter, Invitee: req.UID})
	// 之前已经是好友了，如果是没有使用过邀请码，邀请码还可以给其他人使用；
	// 查找inviteCode对应的邀请人，添加好友
	// 双方的通讯录版本号+1
	// 组装对应的notify信息，通过主服务grpc接口分发消息
	return &types.AddFriendRes{Inviter: invite.Inviter}, nil
}
