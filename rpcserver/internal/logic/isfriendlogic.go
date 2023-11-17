package logic

import (
	"context"

	"github.com/difftim/friend/models"
	"github.com/difftim/friend/rpcserver/internal/svc"
	"github.com/difftim/friend/rpcserver/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFriendLogic {
	return &IsFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFriendLogic) IsFriend(in *pb.IsFriendRequest) (*pb.IsFriendResponse, error) {
	db := l.svcCtx.DbEngine
	base := &pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	// 对 in.CheckList 去重
	var checkList []string = make([]string, 0, len(in.CheckList))
	m := make(map[string]bool)
	for _, v := range in.CheckList {
		if !m[v] {
			m[v] = true
			checkList = append(checkList, v)
		}
	}
	var count int64
	err := db.Model(&models.FriendRelation{}).
		Where("user_id1 = ? and user_id2 in ? or user_id2 = ? and user_id1 in ?",
			in.Uid, checkList, in.Uid, checkList).Count(&count).Error
	if err != nil {
		l.Logger.Errorw("IsFriend failed", logx.Field("err", err))
		base.Reason = err.Error()
		base.Status = 99
		return &pb.IsFriendResponse{Base: base}, nil
	}
	if count < int64(len(checkList)) {
		base.Status = 1
	}

	return &pb.IsFriendResponse{Base: base}, nil
}
