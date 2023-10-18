package logic

import (
	"context"

	commonlogic "chative-server-go/logic"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *pb.AddRequest) (*pb.AddResponse, error) {
	err := commonlogic.AddFriend(
		commonlogic.NewContext(l.ctx, l.Logger, l.svcCtx.DbEngine, l.svcCtx.Redis),
		&commonlogic.AddReq{Inviter: in.Inviter, Invitee: in.Invitee})
	respBase := pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	if err != nil {
		respBase.Status = 19999
		respBase.Reason = err.Error()
		return &pb.AddResponse{Base: &respBase}, err
	}
	return &pb.AddResponse{Base: &respBase}, err
}
