package logic

import (
	"context"

	"chative-server-go/cron"
	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateReminderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateReminderLogic {
	return &CreateReminderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateReminderLogic) CreateReminder(req *types.CreateReminderReq) (resp *types.CreateReminderResp, errInfo *response.ErrInfo) {
	db := l.svcCtx.DbEngine
	var reminder = models.Reminder{
		Version:     1,
		Type:        req.Type,
		Timezone:    req.Timezone,
		Timestamp:   req.Timestamp,
		Repeat:      req.Repeat,
		Description: req.Description,

		Creator: req.UID,
	}
	if reminder.Type == "group" {
		reminder.Conversation = req.Conversation
		//检查是否为群成员
		ok, err := models.IsGroupMember(db, req.Conversation, req.UID)
		if err != nil {
			errInfo = &response.ErrInfo{
				ErrCode: 99,
				Reason:  "Create reminder failed.",
			}
			l.Errorw("Create reminder failed,check member error.", logx.Field("err", err), logx.Field("uid", req.UID))
			return
		}
		if !ok {
			errInfo = &response.ErrInfo{
				ErrCode: 2,
				Reason:  "Not group member.",
			}
			l.Errorw("Create reminder failed,not group member.", logx.Field("uid", req.UID))
			return
		}
	} else {
		if req.Conversation > req.UID {
			reminder.Conversation = req.UID + ":" + req.Conversation
		} else {
			reminder.Conversation = req.Conversation + ":" + req.UID
		}
	}

	err := db.Create(&reminder).Error
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Create reminder failed.",
		}
		l.Errorw("Create reminder failed.", logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}
	cron.NewReminderRun(&reminder)
	resp = &types.CreateReminderResp{
		ID:          reminder.ID,
		Version:     reminder.Version,
		Type:        reminder.Type,
		Timezone:    reminder.Timezone,
		Timestamp:   reminder.Timestamp,
		ModifyTime:  reminder.UpdatedAt.UnixMilli(),
		Repeat:      reminder.Repeat,
		Description: reminder.Description,

		// Creator:     reminder.Creator,
	}

	dat := models.ReminderNotify{
		Operator: req.UID, OperatorDeviceID: req.DID,
		Version: reminder.Version, ChangeType: 1, Creator: reminder.Creator, Type: reminder.Type,
		Conversation: reminder.Conversation, Timezone: reminder.Timezone, Timestamp: reminder.Timestamp,
		Repeat: reminder.Repeat, ReminderID: reminder.ID, Description: reminder.Description,
	}
	cron.SendReminderNotify(&dat, &reminder)
	return
}
