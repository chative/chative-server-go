package logic

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/mainrpc"
	"chative-server-go/models"
	"chative-server-go/response"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateConversationConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateConversationConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateConversationConfigLogic {
	return &UpdateConversationConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateConversationConfigLogic) UpdateConversationConfig(req *types.UpdateConversationReq) (resp *types.Conversation, errInfo *response.ErrInfo) {
	expiryOK := false
	for _, v := range l.svcCtx.Config.MsgExpiryOpts {
		if v == req.MessageExpiry {
			expiryOK = true
			break
		}
	}
	if !expiryOK {
		errInfo = &response.ErrInfo{
			ErrCode: 1,
			Reason:  "Invalid parameter",
		}
		l.Errorw("UpdateConversationConfig invalid message expiry", logx.Field("req", req))
		return
	}

	uids := strings.Split(req.Conversation, ":")
	if len(uids) != 2 || uids[0] > uids[1] {
		errInfo = &response.ErrInfo{
			ErrCode: 1,
			Reason:  "Invalid parameter",
		}
		l.Errorw("UpdateConversationConfig Invalid parameter", logx.Field("conversation", req.Conversation))
		return
	}
	if uids[0] != req.UID && uids[1] != req.UID {
		errInfo = &response.ErrInfo{
			ErrCode: 2,
			Reason:  "No permission",
		}
		l.Errorw("UpdateConversationConfig No permission", logx.Field("conversation", req.Conversation), logx.Field("uid", req.UID))
		return
	}

	//

	db := l.svcCtx.DbEngine
	err := db.Transaction(func(tx *gorm.DB) error {
		// 1. 直接update
		res := tx.Model(models.ShareConversationCnf{}).Where("conversation = ?", req.Conversation).
			Updates(map[string]interface{}{
				"last_operator":     req.UID,
				"last_operator_did": req.DID,
				"message_expiry":    req.MessageExpiry,
				"version":           gorm.Expr("version + 1"),
			}) //map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)}
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			// 2. insert
			res = tx.Create(&models.ShareConversationCnf{
				LastOperator:    req.UID,
				LastOperatorDid: req.DID,
				Conversation:    req.Conversation,
				MessageExpiry:   req.MessageExpiry,
				Version:         1,
			})
			if res.Error != nil {
				return res.Error
			}
		}
		var cnf models.ShareConversationCnf
		resp = &types.Conversation{Conversation: req.Conversation}
		err := tx.Model(&cnf).Where("conversation = ?", req.Conversation).First(&cnf).Error
		if err != nil {
			return err
		}
		resp.MessageExpiry, resp.Ver = cnf.MessageExpiry, cnf.Version
		return nil
	})
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Internal error",
		}
		l.Errorw("UpdateConversationConfig Internal error", logx.Field("err", err))
		return
	}
	// 3. todo 删除缓存
	// 发送notification
	d, _ := json.Marshal(&models.Notify{
		NotifyType: models.DTServerNotifyTypeShareConfig,
		NotifyTime: time.Now().UnixMilli(),
		Data: &models.SharingConversationCnfNotify{
			Operator: req.UID, OperatorDeviceID: req.DID,
			Conversation:  req.Conversation,
			MessageExpiry: resp.MessageExpiry, ChangeType: 1, Ver: resp.Ver},
	})
	err = mainrpc.SendNotify(string(d), uids, "")
	if err != nil {
		l.Errorw("UpdateConversationConfig SendNotify error", logx.Field("err", err))
	}

	return
}
