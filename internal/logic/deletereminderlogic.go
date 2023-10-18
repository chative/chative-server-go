package logic

import (
	"context"
	"strings"

	"chative-server-go/cron"
	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteReminderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteReminderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteReminderLogic {
	return &DeleteReminderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteReminderLogic) DeleteReminder(req *types.DeleteReminderReq) (resp *types.DeleteReminderResp, errInfo *response.ErrInfo) {
	// todo: check permission
	db := l.svcCtx.DbEngine
	var rows []models.Reminder
	err := db.Where("id in ?", strings.Split(req.IDs, ",")).Find(&rows).Error
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Delete reminder failed.",
		}
		l.Errorw("Delete reminder failed, find reminder error.", logx.Field("err", err), logx.Field("uid", req.UID))
		return
	}

	// result := db.Where("id in ?", strings.Split(req.IDs, ",")).Delete(&models.Reminder{})
	// err := result.Error
	// if err != nil {
	// 	errInfo = &response.ErrInfo{
	// 		ErrCode: 99,
	// 		Reason:  "Delete reminder failed.",
	// 	}
	// 	l.Errorw("Delete reminder failed.", logx.Field("err", err), logx.Field("uid", req.UID))
	// 	return
	// }
	// resp = &types.DeleteReminderResp{}
	// if result.RowsAffected == 0 {
	// 	l.Errorw("Delete reminder failed.", logx.Field("err", err), logx.Field("uid", req.UID))
	// 	return
	// }
	// ids := strings.Split(req.IDs, ",")
	// for _, id := range ids {
	// 	rid, _ := strconv.ParseUint(id, 10, 64)
	// 	cron.RemoveReminder(uint(rid))
	// }
	for _, reminder := range rows {
		//检查是否为群成员
		if reminder.Type == "group" {
			ok, err := models.IsGroupMember(db, reminder.Conversation, req.UID)
			if err != nil {
				errInfo = &response.ErrInfo{
					ErrCode: 99,
					Reason:  "System error.",
				}
				l.Errorw("Delete reminder failed,check member error.", logx.Field("err", err), logx.Field("uid", req.UID))
				return
			}
			if !ok {
				errInfo = &response.ErrInfo{
					ErrCode: 2,
					Reason:  "Not group member.",
				}
				l.Errorw("Delete reminder failed,not group member.", logx.Field("uid", req.UID))
				return
			}
		} else {
			if !strings.Contains(reminder.Conversation, req.UID) {
				errInfo = &response.ErrInfo{
					ErrCode: 2,
					Reason:  "Not your conversation",
				}
				l.Errorw("Delete reminder failed,Not your conversation.", logx.Field("uid", req.UID))
				return
			}
		}

		err = db.Where("id = ?", reminder.ID).Delete(&models.Reminder{}).Error
		if err != nil {
			l.Errorw("Delete reminder failed.", logx.Field("err", err),
				logx.Field("uid", req.UID), logx.Field("reminderID", reminder.ID))
			continue
		}
		cron.RemoveReminder(reminder.ID)

		dat := models.ReminderNotify{
			Operator: req.UID, OperatorDeviceID: req.DID,
			Version: reminder.Version, ChangeType: 3, Creator: reminder.Creator, Type: reminder.Type,
			Conversation: reminder.Conversation, Timezone: reminder.Timezone, Timestamp: reminder.Timestamp,
			Repeat: reminder.Repeat, ReminderID: reminder.ID, Description: reminder.Description,
		}

		cron.SendReminderNotify(&dat, &reminder)
	}
	return
}
