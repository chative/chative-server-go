package logic

import (
	"context"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type LinkInvitationInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLinkInvitationInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LinkInvitationInfoLogic {
	return &LinkInvitationInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LinkInvitationInfoLogic) LinkInvitationInfo(req *types.GroupInviteInfoReq) (resp *types.GroupInviteInfoRes, errInfo *response.ErrInfo) {
	inviteInfo, errCode, err := GetInviteInfo(l.svcCtx.Redis, l.svcCtx.DbEngine, req.InviteCode)
	resp = &types.GroupInviteInfoRes{
		GroupName: inviteInfo.Group.Name,
	}
	defer func() {
		if errInfo != nil {
			if errCode == 10120 {
				errInfo.Reason = "Invalid group invite link."
			} else if errCode == 10121 {
				errInfo.Reason = "Failed to join the group. This group disabled invite link."
			} else if errCode == 10122 {
				errInfo.Reason = "Failed to join the group. This group only allows moderators to invite."
			} else if errCode == 10123 {
				errInfo.Reason = "Failed to join the group. This group has already been disbanded."
			} else if errCode == 10124 {
				errInfo.Reason = "Failed to join the group. This group is invalid."
			} else {
				errInfo.Reason = "Internal error"
			}
		}
	}()
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: errCode,
		}
		l.Errorw("LinkInvitationInfo error", logx.Field("err", err))
		return
	}
	// errCode = CheckInvitePermission(inviteInfo)
	// if errCode != 0 {
	// 	errInfo = &response.ErrInfo{
	// 		ErrCode: errCode,
	// 	}
	// 	return
	// }

	userInfo, err := logic.GetUserBasicInfo(l.svcCtx.Redis, l.svcCtx.DbEngine, inviteInfo.GroupMember.Uid)
	if err != nil {
		l.Errorw("LinkInvitationInfo GetUserBasicInfo error", logx.Field("err", err))
		// return
	}
	resp.InviterName = userInfo.PlainName
	resp.AvatarContent, err = FetchAvatar(inviteInfo.Group.Avatar, l.svcCtx.Redis, l.Logger)
	if err != nil {
		l.Errorw("LinkInvitationInfo FetchAvatar error", logx.Field("err", err))
		return
	}
	return
}
