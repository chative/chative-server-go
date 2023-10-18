package logic

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	commonlogic "chative-server-go/logic"
	"chative-server-go/mainrpc"
	"chative-server-go/models"
	"chative-server-go/rediscluster"
	"chative-server-go/response"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendAgreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendAgreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendAgreeLogic {
	return &FriendAgreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendAgreeLogic) FriendAgree(req *types.FriendAgreeReq) (resp *types.FriendAgreeRes, errInfo *response.ErrInfo) {
	db := l.svcCtx.DbEngine

	var ask = models.AskNewFriend{}
	err := db.First(&ask, req.AskID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		errInfo = &response.ErrInfo{
			ErrCode: 19008,
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree ask not found", logx.Field("err", err))
		return
	}
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree db error", logx.Field("err", err))
		return
	}

	isFriend, err := models.IsFriend(db, ask.Inviter, req.Invitee)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree models.IsFriend error", logx.Field("err", err))
		return
	}
	if isFriend {
		errInfo = &response.ErrInfo{
			ErrCode: 19005,
			Reason:  "You are already friends.",
		}
		l.Errorw("FriendAgree already friend", logx.Field("ask", ask))
		return
	}

	var accInviterInfo = models.InternalAccount{Number: ask.Inviter}
	db.Where(accInviterInfo).First(&accInviterInfo)
	if !accInviterInfo.Registered {
		errInfo = &response.ErrInfo{
			ErrCode: 19009,
			Reason:  "This account has logged out and messages can not be reached.",
		}
		l.Errorw("FriendAgree account not registered", logx.Field("ask", ask))
		return
	}

	// timeInerval := time.Duration(l.svcCtx.Config.ReqAddFriend.AskTimeout) * time.Minute
	// timeout := time.Since(ask.CreatedAt) > timeInerval
	timeout := false
	if timeout || ask.Status != 1 {
		errInfo = &response.ErrInfo{
			ErrCode: 19006, // client should tip user to new ask;client特殊处理了
			Reason:  "The request is already expired.",
		}
		l.Errorw("FriendAgree ask status timeout", logx.Field("ask", ask))
		return
	}
	if ask.Status != 1 {
		errInfo = &response.ErrInfo{
			ErrCode: 19007, // client should tip user to new ask;client特殊处理了
			Reason:  "You are not friend yet.",
		}
		l.Errorw("FriendAgree ask status error", logx.Field("ask", ask), logx.Field("timeout", timeout))
		return
	}
	if ask.Invitee != req.Invitee {
		errInfo = &response.ErrInfo{
			ErrCode: 19008, // client should tip user to new ask
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree ask bad request", logx.Field("ask", ask), logx.Field("timeout", timeout))
		return
	}
	//
	//
	logicCtx := logic.NewContext(l.ctx, l.Logger, db, rediscluster.GetRedis())
	blocked, err := logic.InBlockList(logicCtx, ask.Inviter, req.Invitee)
	if blocked {
		// errInfo = &response.ErrInfo{
		// 	ErrCode: 19010,
		// 	Reason:  "You are blocked by this user.",
		// }
		l.Errorw("FriendAgree user in block list", logx.Field("uid", req.Invitee), logx.Field("ask", ask))
		return
	}

	// 添加好友
	err = models.CreateFriendRelation(db, ask.Inviter, req.Invitee)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree models.CreateFriendRelation error", logx.Field("err", err))
		return
	}
	commonlogic.DelFriendCache(l.svcCtx.Redis, ask.Inviter, req.Invitee)
	// 更新ask状态
	go func() {
		err = db.Model(&models.AskNewFriend{}).
			Where(&models.AskNewFriend{Invitee: ask.Invitee, Inviter: ask.Inviter, Status: 1}).
			Or(&models.AskNewFriend{Invitee: ask.Inviter, Inviter: ask.Invitee, Status: 1}).
			Updates(&models.AskNewFriend{Status: 2}).Error
		if err != nil {
			l.Errorw("update ask status error:", logx.Field("err", err))
		}
	}()

	version, err := commonlogic.AddDirectoryVersion(l.svcCtx.Redis, ask.Inviter)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree AddDirectoryVersion error", logx.Field("err", err))
		return
	}
	// 发给对方
	cacheData, err := commonlogic.GetUserBasicInfo(l.svcCtx.Redis, db, ask.Invitee)
	// var accInfo = models.InternalAccount{Number: ask.Invitee}
	// cacheData, err := l.svcCtx.Redis.Get(context.Background(), "Account5"+accInfo.Number).Result()
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree get account info from redis error",
			logx.Field("err", err), logx.Field("uid", ask.Invitee))
		return
	}

	notifyData := &models.AskFriendNotify{AskID: ask.ID, DirectoryVersion: int(version), ActionType: 2} // 2: 同意
	notifyData.OperatorInfo.UID = ask.Invitee
	notifyData.OperatorInfo.Did = req.DID
	notifyData.OperatorInfo.Name = cacheData.PlainName
	notifyData.OperatorInfo.Avatar = cacheData.Avatar2
	notifyData.OperatorInfo.PublicConfigs, _ = json.Marshal(cacheData.PublicConfigs)
	d, _ := json.Marshal(&models.Notify{
		NotifyType: models.DTServerNotifyTypeAddContacts,
		NotifyTime: time.Now().UnixMilli(),
		Data:       notifyData})
	// apn := models.NewApnInfo()
	// apn.SetLocKey("FRIEND_ASK_AGREE").SetLocArgs([]interface{}{}).
	// 	SetLocArgs([]interface{}{cacheData.PlainName, ask.Invitee}).
	// 	SetBody(cacheData.PlainName + " has accepted your request.").
	// 	SetMsg(string(d)).SetPassthrough(`{"conversationId" : "` + ask.Invitee + `" }`)
	// apnData, _ := json.Marshal(apn)
	err = mainrpc.SendNotify(string(d), []string{ask.Inviter}, "")
	if err != nil {
		l.Errorw("FriendAgree SendNotify to Inviter error", logx.Field("err", err))
	}

	// 发给自己
	version, err = commonlogic.AddDirectoryVersion(l.svcCtx.Redis, ask.Invitee)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree AddDirectoryVersion error", logx.Field("err", err))
		return
	}
	accInviterCache, err := commonlogic.GetUserBasicInfo(l.svcCtx.Redis, db, ask.Inviter)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("FriendAgree get account info from redis error",
			logx.Field("err", err), logx.Field("uid", ask.Inviter))
		return
	}
	directoryData := &models.DirectoryNotify{Ver: 1, DirectoryVersion: version}
	member := models.DirectoryMember{ExtID: 2, Number: ask.Inviter,
		Name: accInviterInfo.Name, Action: 0, Avatar: accInviterCache.Avatar2}
	member.PublicConfigs, _ = json.Marshal(accInviterCache.PublicConfigs)
	directoryData.Members = append(directoryData.Members, member)
	d, _ = json.Marshal(&models.Notify{
		NotifyType: models.DTServerNotifyTypeDirectory,
		NotifyTime: time.Now().UnixMilli(),
		Data:       directoryData})
	err = mainrpc.SendNotify(string(d), []string{ask.Invitee}, "")
	if err != nil {
		l.Errorw("FriendAgree SendNotify to Inviter error", logx.Field("err", err))
	}
	return
}
