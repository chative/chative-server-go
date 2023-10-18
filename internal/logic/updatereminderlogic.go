package logic

import (
	"context"

	"chative-server-go/cron"
	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/models"
	"chative-server-go/response"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateReminderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateReminderLogic {
	return &UpdateReminderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateReminderLogic) UpdateReminder(req *types.UpdateReminderReq) (resp *types.UpdateReminderResp, errInfo *response.ErrInfo) {
	db := l.svcCtx.DbEngine
	var reminder = models.Reminder{}
	err := db.First(&reminder, req.ID).Error
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Update reminder failed.",
		}
		l.Errorw("Update reminder failed, find reminder error.", logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}

	if reminder.Creator != req.UID {
		errInfo = &response.ErrInfo{
			ErrCode: 2,
			Reason:  "Not creator.",
		}
		l.Errorw("Update reminder failed, not creator.", logx.Field("uid", req.UID))
		return
	}
	reminder.Description = req.Description
	if reminder.Timestamp != req.Timestamp {
		reminder.NextRun = 0
	}
	reminder.Timestamp = req.Timestamp
	reminder.Timezone = req.Timezone
	reminder.Repeat = req.Repeat
	err = db.Model(&models.Reminder{}).Where("id = ?", reminder.ID).
		Updates(map[string]interface{}{
			"description": reminder.Description,
			"timestamp":   reminder.Timestamp,
			"next_run":    reminder.NextRun,
			"timezone":    reminder.Timezone,
			"repeat":      reminder.Repeat,
			"version":     gorm.Expr("version + 1"),
		}).Error
	// err = db.Save(&reminder).Error
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Update reminder failed.",
		}
		l.Errorw("Update reminder failed, save reminder error.", logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}
	db.First(&reminder, req.ID)
	cron.UpdateReminder(&reminder)
	dat := models.ReminderNotify{
		Operator: req.UID, OperatorDeviceID: req.DID,
		Version: reminder.Version, ChangeType: 2, Creator: reminder.Creator, Type: reminder.Type,
		Conversation: reminder.Conversation, Timezone: reminder.Timezone, Timestamp: reminder.Timestamp,
		Repeat: reminder.Repeat, ReminderID: reminder.ID, Description: reminder.Description,
	}
	cron.SendReminderNotify(&dat, &reminder)

	resp = &types.UpdateReminderResp{
		ID:          reminder.ID,
		Version:     reminder.Version,
		Type:        reminder.Type,
		Timezone:    reminder.Timezone,
		Timestamp:   reminder.Timestamp,
		ModifyTime:  reminder.CreatedAt.UnixMilli(),
		Repeat:      reminder.Repeat,
		Description: reminder.Description,
	}

	return
}
