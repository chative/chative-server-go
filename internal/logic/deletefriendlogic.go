package logic

import (
	"context"
	"encoding/json"
	"time"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	commonlogic "chative-server-go/logic"
	"chative-server-go/mainrpc"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFriendLogic) DeleteFriend(req *types.DelFriendReq) (resp *types.EmptyRes, errInfo *response.ErrInfo) {
	userID1, userID2 := req.Operator, req.UID
	if userID1 > userID2 {
		userID1, userID2 = userID2, userID1
	}
	db := l.svcCtx.DbEngine
	res := db.Delete(&models.FriendRelation{}, &models.FriendRelation{UserID1: userID1, UserID2: userID2})
	err := res.Error
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Delete friend relation failed ",
		}
		l.Errorw("Delete friend relation failed", logx.Field("err", err))
	}
	redisCmd := l.svcCtx.Redis
	commonlogic.DelFriendCache(redisCmd, userID1, userID2)
	if res.RowsAffected == 0 {
		return
	}
	// 发给对方
	version, err := commonlogic.AddDirectoryVersion(redisCmd, req.UID)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("DeleteFriend AddDirectoryVersion error", logx.Field("err", err))
		return
	}
	dat := &models.DirectoryNotify{Ver: 1, DirectoryVersion: version}
	dat.Members = append(dat.Members, models.DirectoryMember{ExtID: 2, Number: req.Operator, Action: 2})
	d, _ := json.Marshal(&models.Notify{
		NotifyType: models.DTServerNotifyTypeDirectory,
		NotifyTime: time.Now().UnixMilli(),
		Data:       dat})
	err = mainrpc.SendNotify(string(d), []string{req.UID}, "")
	if err != nil {
		l.Errorw("DeleteFriend SendNotify to Inviter error", logx.Field("err", err))
	}

	// 发给自己
	version, err = commonlogic.AddDirectoryVersion(redisCmd, req.Operator)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("DeleteFriend AddDirectoryVersion error", logx.Field("err", err))
		return
	}
	directoryData := &models.DirectoryNotify{Ver: 1, DirectoryVersion: version}
	directoryData.Members = append(directoryData.Members, models.DirectoryMember{ExtID: 2, Number: req.UID, Action: 2})
	d, _ = json.Marshal(&models.Notify{
		NotifyType: models.DTServerNotifyTypeDirectory,
		NotifyTime: time.Now().UnixMilli(),
		Data:       directoryData})
	err = mainrpc.SendNotify(string(d), []string{req.Operator}, "")
	if err != nil {
		l.Errorw("DeleteFriend SendNotify to Inviter error", logx.Field("err", err))
	}
	// 删除好友查询path
	db.Delete(&models.UserFindPath{},
		&models.UserFindPath{Src: userID1, Dst: userID2})
	db.Delete(&models.UserFindPath{},
		&models.UserFindPath{Src: userID2, Dst: userID1})
	return
}
