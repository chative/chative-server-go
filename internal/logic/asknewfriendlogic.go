package logic

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	"chative-server-go/mainrpc"
	"chative-server-go/models"
	"chative-server-go/rediscluster"
	"chative-server-go/response"

	"github.com/go-redis/redis_rate/v9"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type AskNewFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAskNewFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AskNewFriendLogic {
	return &AskNewFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AskNewFriendLogic) askNewLimitCheck(req *types.AskFriendReq) bool {
	limiter := redis_rate.NewLimiter(l.svcCtx.Redis)
	res, err := limiter.Allow(l.ctx, strings.Join([]string{"limit:askfriend:single", req.Inviter}, ":"), rediscluster.AskFriendSingleLimit)
	if err != nil {
		l.Errorw("AskNewFriend limit:askfriend:single check error", logx.Field("err", err))
		return false
	}
	if res.Allowed < 1 {
		return false
	}
	userID1, userID2 := req.Inviter, req.Invitee
	if userID1 > userID2 {
		userID1, userID2 = userID2, userID1
	}
	res, err = limiter.Allow(l.ctx, strings.Join([]string{"limit:askfriend:mutual", userID1, userID2}, ":"), rediscluster.AskFriendMutualLimit)
	if err != nil {
		l.Errorw("AskNewFriend limit:askfriend:mutual check error", logx.Field("err", err))
		return false
	}
	if res.Allowed < 1 {
		return false
	}
	return true
}

func (l *AskNewFriendLogic) AskNewFriend(req *types.AskFriendReq) (resp *types.AskFriendRes, errInfo *response.ErrInfo) {
	//  限速
	// if !l.askNewLimitCheck(req) {
	// 	errInfo = &response.ErrInfo{ErrCode: 15, Reason: "You are operating too often, please try later."}
	// 	l.Errorw("too often", logx.Field("req", req))
	// 	return
	// }

	db := l.svcCtx.DbEngine
	// 已经是好友
	isFriend, err := models.IsFriend(db, req.Inviter, req.Invitee)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Internal error",
		}
		l.Errorw("AskNewFriend db error", logx.Field("err", err))
		return
	}
	if isFriend {
		// errInfo = &response.ErrInfo{
		// 	ErrCode: 19004,
		// 	Reason:  "You have already added this friend.",
		// }
		// l.Errorw("AskNewFriend already friend", logx.Field("inviter", req.Inviter), logx.Field("invitee", req.Invitee))
		resp = &types.AskFriendRes{AskID: 0}
		return
	}
	logicCtx := logic.NewContext(l.ctx, l.Logger, db, rediscluster.GetRedis())
	// blocked, err := logic.InBlockList(logicCtx, req.Inviter, req.Invitee)
	// if blocked {
	// 	errInfo = &response.ErrInfo{
	// 		ErrCode: 19011,
	// 		Reason:  "You should UnBlock first.",
	// 	}
	// 	l.Infow("InBlockList Exception", logx.Field("err", err), logx.Field("uid", req.Inviter))
	// 	return
	// }
	// if err != nil {
	// 	errInfo = &response.ErrInfo{
	// 		ErrCode: 19999,
	// 		Reason:  "Unkown Server Error.",
	// 	}
	// 	l.Infow("InBlockList Exception", logx.Field("err", err), logx.Field("uid", req.Inviter))
	// 	return
	// }
	// 被block了
	blockedInviter, err := logic.InBlockList(logicCtx, req.Invitee, req.Inviter)
	if blockedInviter {
		l.Infow("InBlockList", logx.Field("uid", req.Inviter), logx.Field("blocked", req.Invitee))
		resp = &types.AskFriendRes{AskID: 0}
		return
	}
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 19999,
			Reason:  "Unkown Server Error.",
		}
		l.Infow("InBlockList Exception", logx.Field("err", err), logx.Field("uid", req.Inviter))
		return
	}

	var accInfo = models.InternalAccount{Number: req.Invitee}
	db.Where(accInfo).First(&accInfo)
	if !accInfo.Registered {
		errInfo = &response.ErrInfo{
			ErrCode: 19009,
			Reason:  "This account has logged out and messages can not be reached.",
		}
		l.Errorw("FriendAgree account not registered", logx.Field("req", req))
		return
	}

	// 检查同时加好友
	if l.TryAgree(req) {
		resp = &types.AskFriendRes{AskID: -1}
		return
	}

	// 已经有对应的最近请求了
	var recordCreated = false
	var ask = models.AskNewFriend{Inviter: req.Inviter, Invitee: req.Invitee, InviterDid: req.DID}
	err = db.Where(models.AskNewFriend{Inviter: req.Inviter, Invitee: req.Invitee}).
		Order("id desc").First(&ask).Error
	if err == nil && ask.Status == 1 {
		// errInfo = &response.ErrInfo{
		// 	ErrCode: 1,
		// 	Reason:  "Already asked",
		// }
		// l.Errorw("AskNewFriend already asked", logx.Field("req", req))
		// err = sendAskNewFriendNotify(db, req, ask.ID)
		// return
		// err = db.Model(ask).Update("created_at", gorm.Expr("now()")).Error
		l.Infow("AskNewFriend already asked", logx.Field("req", req), logx.Field("ask", ask))
	} else {
		// db.Delete(&ask)  // 先不删除无效数据
		ask = models.AskNewFriend{Inviter: req.Inviter, Invitee: req.Invitee, InviterDid: req.DID, Status: 1}
		err = db.Save(&ask).Error
		if err != nil {
			errInfo = &response.ErrInfo{
				ErrCode: 19999,
				Reason:  "Internal error",
			}
			l.Errorw("AskNewFriend Create askRecord error", logx.Field("err", err))
			return
		}
		recordCreated = true
	}
	go func() {
		// 保存find path
		if recordCreated && req.Source != nil {
			err = logic.SaveFindPath(logicCtx, req.Inviter, req.Invitee, models.FindPath{
				Type:    req.Source.Type,
				GroupID: req.Source.GroupID,
				UID:     req.Source.UID,
			})
			if err != nil {
				l.Errorw("AskNewFriend SaveFindPath error", logx.Field("err", err))
			}
		}
		// 发送请求通知
		err = sendAskNewFriendNotify(db, req, ask.ID)
		if err != nil {
			l.Errorw("AskNewFriend send notify error", logx.Field("err", err))
		}
	}()
	resp = &types.AskFriendRes{AskID: int(ask.ID)}
	return
}

func (l *AskNewFriendLogic) TryAgree(req *types.AskFriendReq) bool {
	db := l.svcCtx.DbEngine
	ask, err := models.FirstValidAsk(db, req.Invitee, req.Inviter)
	if err != nil && req.Action != "accept" { // 非同意操作
		return false
	}
	if ask.ID == 0 {
		err = db.Save(ask).Error
	}
	if err != nil {
		l.Errorw("FriendAgree FirstValidAsk or load error", logx.Field("err", err))
		return false
	}

	_, errInfo := NewFriendAgreeLogic(l.ctx, l.svcCtx).FriendAgree(
		&types.FriendAgreeReq{Invitee: req.Inviter, AskID: ask.ID})
	if errInfo == nil || errInfo.ErrCode == 19005 { // 同意成功 或者 已经是好友了
		return true
	}
	return false
}

func sendAskNewFriendNotify(db *gorm.DB, req *types.AskFriendReq, askID uint) error {
	// notifyData := &models.AskFriendNotify{AskID: askID, DirectoryVersion: int(version), ActionType: 2} // 2: 同意
	// notifyData.OperatorInfo.UID = req.Invitee
	// notifyData.OperatorInfo.Did = req.DID
	// notifyData.OperatorInfo.Name = cacheData.PlainName
	// notifyData.OperatorInfo.Avatar = cacheData.Avatar2
	// notifyData.OperatorInfo.PublicConfigs, _ = json.Marshal(cacheData.PublicConfigs)
	// d, _ := json.Marshal(&models.Notify{
	// 	NotifyType: models.DTServerNotifyTypeAddContacts,
	// 	NotifyTime: time.Now().UnixMilli(),
	// 	Data:       notifyData})

	// err = mainrpc.SendNotify(string(d), []string{req.Inviter},"")
	// if err != nil {
	// 	l.Errorw("FriendAgree SendNotify to Inviter error", logx.Field("err", err))
	// }//
	var accInfo = models.InternalAccount{Number: req.Inviter}
	db.Where(accInfo).First(&accInfo)
	notifyData := &models.AskFriendNotify{AskID: askID, ActionType: 1}
	notifyData.OperatorInfo.UID = req.Inviter
	notifyData.OperatorInfo.Did = req.DID
	notifyData.OperatorInfo.Name = accInfo.Name
	d, _ := json.Marshal(&models.Notify{
		NotifyType: models.DTServerNotifyTypeAddContacts,
		NotifyTime: time.Now().UnixMilli(),
		Data:       notifyData})
	// models.NewAppInfo()
	// apn := models.NewApnInfo()
	// apn.SetLocKey("FRIEND_ASK_NEW").SetLocArgs([]interface{}{}).
	// 	SetLocArgs([]interface{}{accInfo.Name, askID}).
	// 	SetBody(accInfo.Name + " sent you a contact request.").
	// 	SetMsg(string(d)).SetPassthrough(`{"conversationId" : "` + req.Inviter + `" }`)
	// apnData, _ := json.Marshal(apn)
	err := mainrpc.SendNotify(string(d), []string{req.Invitee}, "")
	return err
}
