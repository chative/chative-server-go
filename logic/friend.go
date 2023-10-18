package logic

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"chative-server-go/mainrpc"
	"chative-server-go/models"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// var friendService FriendService

// type FriendService struct {
// }

type Context struct {
	logx.Logger
	ctx context.Context
	db  *gorm.DB

	redisCmd redis.Cmdable
}

func NewContext(ctx context.Context, logger logx.Logger,
	db *gorm.DB, redisCmd redis.Cmdable) *Context {
	return &Context{
		Logger:   logger,
		ctx:      ctx,
		db:       db,
		redisCmd: redisCmd,
	}
}

func InBlockList(ctx *Context, userA, userB string) (bool, error) {
	db := ctx.db
	return models.ExistConversation(db, userA, userB, "1")
}

func AddFriend(ctx *Context, req *AddReq) error {
	//
	// 查看是否为好友
	var friendRelation = &models.FriendRelation{UserID1: req.Inviter, UserID2: req.Invitee}
	if friendRelation.UserID1 > friendRelation.UserID2 {
		friendRelation.UserID1, friendRelation.UserID2 = friendRelation.UserID2, friendRelation.UserID1
	}
	db := ctx.db
	tNow := time.Now()
	err := db.FirstOrCreate(friendRelation, friendRelation).Error
	DelFriendCache(ctx.redisCmd, friendRelation.UserID1, friendRelation.UserID2)
	if err != nil {
		ctx.Errorw("db.FirstOrCreate failed", logx.Field("err", err))
		return err
	}
	if friendRelation.UpdatedAt.UnixMilli() < tNow.UnixMilli() {
		ctx.Infow("AddFriend", logx.Field("friendRelation", friendRelation))
		return nil
	}
	// 发送通知
	addContactVersion(ctx, req.Inviter, req.Invitee, ctx.redisCmd)
	addContactVersion(ctx, req.Invitee, req.Inviter, ctx.redisCmd)
	ctx.Infow("AddFriend", logx.Field("req", req))
	return nil
}

func AddDirectoryVersion(redisCmd redis.Cmdable, uid string) (int64, error) {
	return redisCmd.Incr(context.Background(), "PersonalDirectoryVersion_"+uid).Result()
}

func DelFriendCache(edisCmd redis.Cmdable, userID1, userID2 string) {
	if userID1 > userID2 {
		userID1, userID2 = userID2, userID1
	}
	key := strings.Join([]string{"friend", userID1, userID2}, "_")
	edisCmd.Del(context.Background(), key)
}

func addContactVersion(ctx *Context, to, uid string, redisCmd redis.Cmdable) (err error) {
	// to 接收notify的人
	tNow := time.Now()
	notify := models.Notify{
		NotifyType: models.DTServerNotifyTypeDirectory, NotifyTime: tNow.UnixMilli(),
	}
	dat := models.DirectoryNotify{}
	dat.Ver = 1
	dat.DirectoryVersion, err = AddDirectoryVersion(redisCmd, to) //redisCmd.Incr(context.Background(), "PersonalDirectoryVersion_"+to).Result()
	if err != nil {
		ctx.Errorw("redisCmd.Incr failed", logx.Field("err", err), logx.Field("to", to))
	}
	cacheData, err := GetUserBasicInfo(ctx.redisCmd, ctx.db, uid)
	if err != nil {
		ctx.Errorw("GetUserCache failed", logx.Field("err", err), logx.Field("uid", uid))
	}
	member := models.DirectoryMember{ExtID: 2, Number: uid, Name: cacheData.PlainName, Action: 0,
		Avatar: cacheData.Avatar2}
	if member.Name == "" {
		member.Name = uid
	}
	member.PublicConfigs, _ = json.Marshal(cacheData.PublicConfigs)
	dat.Members = append(dat.Members, member)
	notify.Data = &dat
	content, _ := json.Marshal(&notify)
	err = mainrpc.SendNotify(string(content), []string{to}, "")
	if err != nil {
		ctx.Errorw("mainrpc.ContactNotify failed", logx.Field("err", err), logx.Field("to", to))
	}
	return err
}

// func ExistFriend(ctx *Context, req *ExistReq) (bool, error) {
// 	// 查看是否为好友
// 	var friendRelation = &models.FriendRelation{UserID1: req.UserID1, UserID2: req.UserID2}
// 	if friendRelation.UserID1 > friendRelation.UserID2 {
// 		friendRelation.UserID1, friendRelation.UserID2 = friendRelation.UserID2, friendRelation.UserID1
// 	}
// 	db := ctx.db
// 	err := db.First(friendRelation, friendRelation).Error
// 	if err != nil {
// 		return false, err
// 	}
// 	ctx.Infow("ExistFriend", logx.Field("friendRelation", friendRelation))
// 	return true, nil
// }

func ListFriends(ctx *Context, req *ListReq) ([]string, error) {
	//
	var friendRelation = &models.FriendRelation{UserID1: req.UserID}
	db := ctx.db
	var friendRelations []*models.FriendRelation
	err := db.Where(friendRelation).Or(models.FriendRelation{UserID2: req.UserID}).Find(&friendRelations).Error
	if err != nil {
		return nil, err
	}
	ctx.Infow("ListFriends", logx.Field("friendRelations", friendRelations))
	friends := make([]string, 0, len(friendRelations)+1)
	for _, friendRelation := range friendRelations {
		if friendRelation.UserID1 == req.UserID {
			friends = append(friends, friendRelation.UserID2)
		} else {
			friends = append(friends, friendRelation.UserID1)
		}
	}
	friends = append(friends, req.UserID)
	return friends, nil

}

// 保存find path
func SaveFindPath(ctx *Context, src, dst string, findPath models.FindPath) error {
	db := ctx.db
	var findPathDB = models.UserFindPath{Src: src,
		Dst: dst, Path: findPath}
	err := db.Save(&findPathDB).Error
	if err != nil {
		err = db.Where(models.UserFindPath{Src: src, Dst: dst}).
			Updates(&findPathDB).Error
	}
	return err
}

func CheckFindPath(ctx *Context, src string, dst []string, findPathType string) {
	chDone := make(chan struct{})
	go func() {
		for _, uid := range dst {
			// 1. 是不是好友
			if ok, err := models.IsFriend(ctx.db, uid, src); ok {
				continue
			} else if err != nil {
				ctx.Logger.Errorw("CheckFindPath IsFriend error", logx.Field("err", err))
			}
			// 2. 有没有有效的好友请求
			err := ctx.db.First(&models.AskNewFriend{}, models.AskNewFriend{Invitee: uid, Inviter: src, Status: 1}).Error
			if err == nil {
				continue
			}
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 3. 保存查找来源
				err = SaveFindPath(ctx, src, uid, models.FindPath{Type: findPathType})
				if err != nil {
					ctx.Logger.Errorw("CheckFindPath SaveFindPath error", logx.Field("err", err))
				}
				continue
			}
			if err != nil {
				ctx.Logger.Errorw("CheckFindPath First error", logx.Field("err", err))
			}
		}
		chDone <- struct{}{}
	}()
	select {
	case <-chDone:
	case <-time.After(500 * time.Millisecond):
	}
}
