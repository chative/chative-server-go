package logic

import (
	"context"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/mainrpc"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReportAccLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReportAccLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReportAccLogic {
	return &ReportAccLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReportAccLogic) ReportAcc(req *types.ReportAccReq) (resp *types.ReportAccRes, errInfo *response.ErrInfo) {
	// 创建失败就更新
	db := l.svcCtx.DbEngine
	var reportAcc = models.ReportLog{Informer: req.Informer, Suspect: req.Suspect,
		Reason: req.Reason, Type: req.Type, Block: req.Block}
	err := db.Create(&reportAcc).Error
	if err != nil {
		result := db.Where(models.ReportLog{Informer: req.Informer, Suspect: req.Suspect}).
			Updates(reportAcc)
		if result.Error != nil {
			l.Errorw("ReportAcc failed", logx.Field("req", req), logx.Field("err", err))
			errInfo = &response.ErrInfo{
				ErrCode: 99,
				Reason:  "internal error",
			}
			return
		}
	}
	// 调用block
	if req.Block >= 0 {
		mainrpc.BlockConversation(req.Informer, req.Suspect, int32(req.Block))
	}
	resp = &types.ReportAccRes{}

	return
}
