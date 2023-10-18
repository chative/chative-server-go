package logic

import (
	"context"

	"chative-server-go/logic"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HowToMetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHowToMetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HowToMetLogic {
	return &HowToMetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HowToMetLogic) HowToMet(in *pb.HowToMetRequest) (*pb.HowToMetResponse, error) {
	respBase := pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	var desArr = make([]string, 0, len(in.Src))
	var findyouArr = make([]string, 0, len(in.Src))
	for _, v := range in.Src {
		des, err := logic.FindPathFormat(
			logic.NewContext(l.ctx, l.Logger, l.svcCtx.DbEngine, l.svcCtx.Redis),
			v, in.Dst, in.Lang, in.SourceQueryType)
		if err != nil {
			l.Errorw("HowToMet failed", logx.Field("src", in.Src), logx.Field("dst", in.Dst),
				logx.Field("err", err))
			respBase.Status = 1
			respBase.Reason = err.Error()
			return &pb.HowToMetResponse{Base: &respBase}, nil
		}
		desArr = append(desArr, des)
		findyou, err := logic.FindPathFormat(
			logic.NewContext(l.ctx, l.Logger, l.svcCtx.DbEngine, l.svcCtx.Redis),
			v, in.Dst, in.Lang, "findyou")
		if err != nil {
			l.Errorw("HowToMet failed", logx.Field("src", in.Src), logx.Field("dst", in.Dst), logx.Field("err", err))
		}
		findyouArr = append(findyouArr, findyou)
	}
	return &pb.HowToMetResponse{Base: &respBase, Describe: desArr, Findyou: findyouArr}, nil
}
