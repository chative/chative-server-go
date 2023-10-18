package logic

import (
	"context"
	"strings"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/response"
	"chative-server-go/utils/iplocation"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLocationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLocationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLocationLogic {
	return &GetLocationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLocationLogic) GetLocation(req *types.LocationReq) (resp *types.LocationRes, errInfo *response.ErrInfo) {
	ips := strings.Split(req.XForwardedFor, ",")
	if len(ips) == 0 {
		errInfo = &response.ErrInfo{
			ErrCode: 1,
			Reason:  "error request",
		}
	}
	loc, err := iplocation.GetLocationPhone(ips[0])
	if err != nil {
		l.Errorw("GetLocationPhone failed", logx.Field("err", err), logx.Field("ip", ips[0]))
		resp = &types.LocationRes{
			CountryName: "United States",
			CountryCode: "US",
			DialingCode: "+1",
		}
		return
	}
	resp = &types.LocationRes{
		CountryName: loc.CountryName,
		CountryCode: loc.CountryCode,
		DialingCode: loc.DialingCode,
	}
	return
}
