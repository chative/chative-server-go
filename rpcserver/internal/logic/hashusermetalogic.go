package logic

import (
	"context"

	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"
	"chative-server-go/utils/crypto"
	"chative-server-go/utils/secretsmanager"

	"github.com/zeromicro/go-zero/core/logx"
)

type HashUserMetaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHashUserMetaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HashUserMetaLogic {
	return &HashUserMetaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HashUserMetaLogic) HashUserMeta(in *pb.HashUserMetaRequest) (*pb.HashUserMetaResponse, error) {
	base := &pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	sm := secretsmanager.GetSM()
	resp := &pb.HashUserMetaResponse{Base: base}
	if in.Email != "" {
		emailHash, err := crypto.HashID(ClearEmail(in.Email), sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("HashUserMeta Failed", logx.Field("err", err))
		}
		resp.EmailHash = emailHash
	}
	if in.Phone != "" {
		phoneHash, err := crypto.HashID(in.Phone, sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("HashUserMeta Failed", logx.Field("err", err))
		}
		resp.PhoneHash = phoneHash
	}
	l.Infow("HashUserMeta", logx.Field("resp", resp))
	return resp, nil
}
