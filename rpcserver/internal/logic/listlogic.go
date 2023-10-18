package logic

import (
	"context"

	commonlogic "chative-server-go/logic"

	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *pb.ListRequest) (*pb.ListResponse, error) {
	friends, err := commonlogic.ListFriends(
		commonlogic.NewContext(l.ctx, l.Logger, l.svcCtx.DbEngine, l.svcCtx.Redis),
		&commonlogic.ListReq{UserID: in.User})
	respBase := pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	if err != nil {
		respBase.Status = 19999
		respBase.Reason = err.Error()
		return &pb.ListResponse{Base: &respBase}, err
	}
	return &pb.ListResponse{Base: &respBase, Friends: friends}, nil
}
