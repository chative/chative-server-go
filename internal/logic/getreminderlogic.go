package logic

import (
	"context"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetReminderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetReminderLogic {
	return &GetReminderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetReminderLogic) GetReminder(req *types.GetReminderReq) (resp *types.GetReminderResp, errInfo *response.ErrInfo) {
	db := l.svcCtx.DbEngine
	var conID string
	if req.Type == "group" {
		conID = req.Conversation
		//检查是否为群成员
		ok, err := models.IsGroupMember(db, req.Conversation, req.UID)
		if err != nil {
			errInfo = &response.ErrInfo{
				ErrCode: 99,
				Reason:  "System error.",
			}
			l.Errorw("get reminder failed,check member error.", logx.Field("err", err), logx.Field("uid", req.UID))
			return
		}
		if !ok {
			errInfo = &response.ErrInfo{
				ErrCode: 2,
				Reason:  "Not group member.",
			}
			l.Errorw("get reminder failed,not group member.", logx.Field("uid", req.UID))
			return
		}
	} else {
		if req.Conversation > req.UID {
			conID = req.UID + ":" + req.Conversation
		} else {
			conID = req.Conversation + ":" + req.UID
		}
	}

	var reminders []models.Reminder
	err := db.Where("conversation = ?", conID).Find(&reminders).Error
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Get reminder failed.",
		}
		l.Errorw("Get reminder failed.", logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}
	resp = &types.GetReminderResp{
		List: make([]types.CreateReminderResp, 0, len(reminders)),
	}
	for _, reminder := range reminders {
		resp.List = append(resp.List, types.CreateReminderResp{
			ID:           reminder.ID,
			Type:         reminder.Type,
			Timezone:     reminder.Timezone,
			Timestamp:    reminder.Timestamp,
			ModifyTime:   reminder.UpdatedAt.UnixMilli(),
			Repeat:       reminder.Repeat,
			Description:  reminder.Description,
			Creator:      reminder.Creator,
			Conversation: req.Conversation,
		})
	}

	return
}
