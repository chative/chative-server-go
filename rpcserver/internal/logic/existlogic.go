package logic

import (
	"context"
	"errors"
	"strings"
	"time"

	"chative-server-go/models"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

const existCacheTimeout = time.Minute * 10

type ExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistLogic {
	return &ExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExistLogic) Exist(in *pb.ExistRequest) (*pb.ExistResponse, error) {
	db := l.svcCtx.DbEngine
	respBase := pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	// 查看是否为好友
	var friendRelation = &models.FriendRelation{UserID1: in.User, UserID2: in.Friend}
	if friendRelation.UserID1 > friendRelation.UserID2 {
		friendRelation.UserID1, friendRelation.UserID2 = friendRelation.UserID2, friendRelation.UserID1
	}
	key := strings.Join([]string{"friend", friendRelation.UserID1, friendRelation.UserID2}, "_")
	red := l.svcCtx.Redis
	s, err := red.GetEx(context.Background(), key, existCacheTimeout).Result()
	l.Infow("FriendExist redis get", logx.Field("key", key), logx.Field("value", s))
	if s == "1" {
		return &pb.ExistResponse{Base: &respBase, Exist: true}, nil
	} else if s == "0" {
		return &pb.ExistResponse{Base: &respBase, Exist: false}, nil
	}
	if err != nil {
		l.Errorw("redis get error", logx.Field("err", err), logx.Field("redisNil", err == redis.Nil))
	}
	err = db.Where(friendRelation).First(friendRelation).Error
	l.Infow("FriendExist db get", logx.Field("friendRelation", friendRelation), logx.Field("err", err))
	if err == nil {
		red.SetEX(context.Background(), key, "1", existCacheTimeout)
		return &pb.ExistResponse{Base: &respBase, Exist: true}, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		red.SetEX(context.Background(), key, "0", existCacheTimeout)
		return &pb.ExistResponse{Base: &respBase, Exist: false}, nil
	} else {
		respBase.Status = 19999
		respBase.Reason = err.Error()
		return &pb.ExistResponse{Base: &respBase, Exist: false}, err
	}
}
