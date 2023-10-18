package logic

import (
	"context"

	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAccountLogic {
	return &DeleteAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAccountLogic) DeleteAccount(in *pb.DeleteAccReq) (*pb.DeleteAccResp, error) {
	// 更新internal account
	db := l.svcCtx.DbEngine
	err := db.Where(&models.InternalAccount{Number: in.Uid}).
		Save(&models.InternalAccount{Number: in.Uid, Deleted: true}).Error
	if err != nil {
		l.Errorw("DeleteInternalAccount Failed", logx.Field("err", err), logx.Field("uid", in.Uid))
		return nil, err
	}
	//
	err = db.Where(&models.Account{Number: in.Uid}).Delete(&models.Account{}).Error
	if err != nil {
		l.Errorw("DeleteAccount Failed", logx.Field("err", err), logx.Field("uid", in.Uid))
		return nil, err
	}
	// conversations
	err = db.Where(&models.Conversation{UID: in.Uid}).Delete(&models.Conversation{}).Error
	if err != nil {
		l.Errorw("DeleteConversation Failed", logx.Field("err", err), logx.Field("uid", in.Uid))
	}
	profile := models.UserProfile{UID: in.Uid}
	db.First(&profile, profile)
	// delete profile、account
	err = db.Model(&models.UserProfile{}).Where(&models.UserProfile{UID: in.Uid}).
		Delete(&models.UserProfile{}).Error
	if err != nil {
		l.Errorw("DeleteUserProfile Failed", logx.Field("err", err), logx.Field("uid", in.Uid))
	}
	// search
	var hKey = logic.DirectoryHashKey
	redisCmd := l.svcCtx.Redis
	if profile.EmailHash != "" {
		redisCmd.HDel(l.ctx, hKey, profile.EmailHash)
	}
	if profile.PhoneHash != "" {
		redisCmd.HDel(l.ctx, hKey, profile.PhoneHash)
	}
	// 清理好友
	err = db.Where(&models.FriendRelation{UserID1: in.Uid}).
		Or(&models.FriendRelation{UserID2: in.Uid}).Delete(&models.FriendRelation{}).Error
	if err != nil {
		l.Errorw("DeleteFriendRelation Failed", logx.Field("err", err), logx.Field("uid", in.Uid))
	}
	respBase := pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	return &pb.DeleteAccResp{Base: &respBase}, nil
}
