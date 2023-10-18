package cron

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"chative-server-go/mainrpc"
	"chative-server-go/models"

	"github.com/go-co-op/gocron"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// 1. 计算下次提醒时间

//
const monitorChannelName = "reminder:monitor:channel"

//
var reminderMgr *ReminderMgr

func SendReminderNotify(dat *models.ReminderNotify, reminder *models.Reminder) (err error) {
	tNow := time.Now()
	notify := models.Notify{
		NotifyType: models.DTServerNotifyTypeReminder, NotifyTime: tNow.UnixMilli(),
	}
	dat.ModifyTime = reminder.UpdatedAt.UnixMilli()
	// 群组提醒
	notify.Data = &dat
	var apn *models.ApnInfo
	switch dat.ChangeType {
	// case 1:
	// 	apn.SetLocKey("REMINDER_CREATED")
	// 	apn.SetBody("You have a new reminder.")
	// case 2:
	// 	apn.SetLocKey("REMINDER_UPDATED")
	// 	apn.SetBody("You have a reminder updated.")
	// case 3:
	// 	apn.SetLocKey("REMINDER_DELETED")
	// 	apn.SetBody("You have a reminder deleted.")
	case 4:
		apn = models.NewApnInfo()
		apn.SetLocKey("REMINDER_REMIND")
		apn.SetBody("Reminder: " + reminder.Description)
		apn.SetLocArgs([]interface{}{reminder.ID})
	}
	if reminder.Type == "group" {
		content, _ := json.Marshal(&notify)
		var apnStr string
		if apn != nil {
			apn.SetMsg(string(content)).SetPassthrough(`{"conversationId" : "` +
				base64.RawStdEncoding.EncodeToString([]byte(reminder.Conversation)) + `" }`)
			apnData, _ := json.Marshal(apn)
			apnStr = string(apnData)
		}
		err = mainrpc.SendGroupNotify(string(content), []string{reminder.Conversation}, apnStr)
		if err != nil {
			logx.Errorw("mainrpc.SendGroupNotify failed",
				logx.Field("err", err), logx.Field("reminderId", reminder.ID))
		}
	} else {
		// private提醒
		uids := strings.Split(reminder.Conversation, ":")
		for i, v := range uids {
			dat.Conversation = strings.Replace(reminder.Conversation, v, "", 1)
			dat.Conversation = strings.Replace(dat.Conversation, ":", "", 1)
			content, _ := json.Marshal(&notify)
			var apnStr string
			if apn != nil {
				apn.SetMsg(string(content)).SetPassthrough(`{"conversationId" : "` +
					dat.Conversation + `" }`)
				apnData, _ := json.Marshal(apn)
				apnStr = string(apnData)
			}
			err = mainrpc.SendNotify(string(content), []string{v}, apnStr)
			if err != nil {
				logx.Errorw("mainrpc.SendNotify failed",
					logx.Field("err", err), logx.Field("reminderId", reminder.ID))
			}
			if i == 0 && len(uids) == 2 && v == uids[1] {
				break
			}
		}
	}
	return
}

func NewReminderRun(reminder *models.Reminder) {
	err := reminderMgr.scheduleOneTask(reminder)
	if err != nil {
		logx.Errorw("in NewReminderRun, scheduleOneTask failed",
			logx.Field("err", err), logx.Field("reminderID", reminder.ID))
	} else {
		reminderMgr.Publish("add", reminder.ID)
	}
}

func UpdateReminder(reminder *models.Reminder) {
	reminderMgr.delTask(reminder.ID)
	err := reminderMgr.scheduleOneTask(reminder)
	if err != nil {
		logx.Errorw("in UpdateReminder, scheduleOneTask failed", logx.Field("reminderID", reminder.ID))
	} else {
		reminderMgr.Publish("update", reminder.ID)
	}
}

func RemoveReminder(reminderID uint) {
	reminderMgr.delTask(reminderID)
	reminderMgr.Publish("remove", reminderID)
}

func doCronTask(reminderID uint) {
	reminderMgr.reminders <- reminderID
}

func initReminder(redisCmd *redis.ClusterClient, db *gorm.DB) {
	reminderMgr = &ReminderMgr{
		scheduler: gocron.NewScheduler(time.UTC),
		redisCmd:  redisCmd,
		db:        db,
		reminders: make(chan uint, 1000),
	}
	reminderMgr.scheduler.TagsUnique()
	reminderMgr.Start()
}

type ReminderMgr struct {
	scheduler *gocron.Scheduler
	db        *gorm.DB
	redisCmd  *redis.ClusterClient
	reminders chan uint
	lock      sync.Mutex
}

func (r *ReminderMgr) Start() {
	rows, err := r.db.Model(&models.Reminder{}).Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var reminder models.Reminder
		err = r.db.ScanRows(rows, &reminder)
		if err != nil {
			panic(err)
		}
		err = r.scheduleOneTask(&reminder)
		if err != nil {
			logx.Errorw("scheduleOneTask failed", logx.Field("reminderID", reminder.ID),
				logx.Field("err", err))
			log.Println("scheduleOneTask failed,reminderID:", reminder.ID, err)
			panic(err)
		}
	}
	go r.runScheduleTask()
	go r.SubscribeChange()
	r.scheduler.StartAsync()
}

func (r *ReminderMgr) scheduleOneTask(reminder *models.Reminder) (err error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	// todo 计算下次提醒时间
	var at time.Time
	if reminder.NextRun > reminder.Timestamp {
		at = time.UnixMilli(reminder.NextRun).UTC()
	} else {
		at = time.UnixMilli(reminder.Timestamp).UTC()
	}
	s := r.scheduler.Tag(strconv.FormatUint(uint64(reminder.ID), 10))
	switch reminder.Repeat {
	case models.ReminderRepeatDaily:
		_, err = s.Every(1).Day().StartAt(at).At(at).Do(doCronTask, reminder.ID)
	case models.ReminderRepeatWeekly:
		_, err = addWeekday(s.Every(1).Week(), at).StartAt(at).At(at).Do(doCronTask, reminder.ID)
	case models.ReminderRepeatBiweekly:
		_, err = addWeekday(s.Every(2).Week(), at).StartAt(at).At(at).Do(doCronTask, reminder.ID)
	case models.ReminderWeekdays:
		_, err = s.Every(1).Week().Monday().Tuesday().Wednesday().Thursday().Friday().
			StartAt(at).At(at).Do(doCronTask, reminder.ID)
	case models.ReminderRepeatMonthly:
		scheduleTime := time.UnixMilli(reminder.Timestamp).UTC()
		day := scheduleTime.Day()
		if day < 29 {
			_, err = s.Every(1).Month(day).StartAt(at).At(at).Do(doCronTask, reminder.ID)
			if err != nil {
				logx.Errorw("in scheduleOneTask, Every Month failed",
					logx.Field("err", err), logx.Field("reminderId", reminder.ID),
					logx.Field("day", day), logx.Field("at", at))
			}
		} else {
			var day int = -1
			switch scheduleTime.Month() {
			case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
				day = scheduleTime.Day() - 1 - 31
			case time.April, time.June, time.September, time.November:
				day = scheduleTime.Day() - 1 - 30
				// case time.February: // 闰年, 29天,就是最后一天
				// 	if isLeap(at.Year()) {

				// 	}
			}
			_, err = s.Every(1).Month(day).StartAt(at).At(at).Do(doCronTask, reminder.ID)
			if err != nil {
				logx.Errorw("in scheduleOneTask, Every Month failed",
					logx.Field("err", err), logx.Field("reminderId", reminder.ID), logx.Field("day", day))
			}
		}
	default:
		_, err = s.Every(1000).Week().StartAt(at).At(at).Do(doCronTask, reminder.ID)
	}
	// r.scheduler.Every(1).Day().At(reminder.Time).Do(doCronTask, reminder.ID)
	return
}

func addWeekday(week *gocron.Scheduler, at time.Time) *gocron.Scheduler {
	switch at.Weekday() {
	case time.Monday:
		week.Monday()
	case time.Tuesday:
		week.Tuesday()
	case time.Wednesday:
		week.Wednesday()
	case time.Thursday:
		week.Thursday()
	case time.Friday:
		week.Friday()
	case time.Saturday:
		week.Saturday()
	case time.Sunday:
		week.Sunday()
	}
	return week
}
func (r *ReminderMgr) runScheduleTask() {
	for id := range r.reminders {
		set, err := r.redisCmd.SetNX(context.Background(), "reminder:exclusive:"+strconv.FormatUint(uint64(id),
			10), 1, 10*time.Second).Result()
		if err != nil {
			logx.Errorw("redisCmd.SetNX failed", logx.Field("err", err), logx.Field("reminderId", id))
		} else if !set {
			logx.Infow("doing on other mechine,skip doing", logx.Field("id", id))
			continue
		}
		r.sendTip(id)
	}
}

func (r *ReminderMgr) delTask(id uint) {
	r.scheduler.RemoveByTag(strconv.FormatUint(uint64(id), 10))
}

func (r *ReminderMgr) updatedNextRun(id uint, job *gocron.Job) {
	r.db.Model(&models.Reminder{}).Where("id = ?", id).Update("next_run", job.NextRun().UnixMilli())
}

func (r *ReminderMgr) sendTip(id uint) (err error) {
	var reminder models.Reminder
	err = r.db.First(&reminder, id).Error
	logx.Infow("in sendTip", logx.Field("reminderId", id), logx.Field("err", err))
	if err != nil {
		logx.Errorw("in sendTip, first reminder, failed", logx.Field("reminderId", id), logx.Field("err", err))
		return
	}
	if reminder.Repeat == 0 {
		r.db.Delete(&models.Reminder{}, id)
		r.scheduler.RemoveByTag(strconv.FormatUint(uint64(id), 10))
	} else {
		jobs, err := r.scheduler.FindJobsByTag(strconv.FormatUint(uint64(id), 10))
		if err != nil {
			logx.Errorw("FindJobsByTag failed", logx.Field("reminderId", id), logx.Field("err", err))
		} else {
			r.updatedNextRun(id, jobs[0])
		}
	}

	dat := models.ReminderNotify{
		Version: reminder.Version, ChangeType: 4, Creator: reminder.Creator, Type: reminder.Type,
		Conversation: reminder.Conversation, Timezone: reminder.Timezone, Timestamp: reminder.Timestamp,
		Repeat: reminder.Repeat, ReminderID: reminder.ID, Description: reminder.Description,
	}
	SendReminderNotify(&dat, &reminder)
	return
}

type channelPayload struct {
	Op string `json:"op"`
	ID uint   `json:"id"`
}

// 订阅 redis channel "reminder:monitor:channel"
func (r *ReminderMgr) SubscribeChange() {
	pubsub := r.redisCmd.Subscribe(context.Background(), monitorChannelName)
	for {
		func() {
			defer func() {
				err := recover()
				if err != nil {
					logx.Errorw("in SubscribeChange, panic", logx.Field("err", err))
				}
			}()
			msg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				logx.Errorw("in SubscribeChange, ReceiveMessage failed", logx.Field("err", err))
				return
			}
			var payload channelPayload
			err = json.Unmarshal([]byte(msg.Payload), &payload)
			if err != nil {
				logx.Errorw("in SubscribeChange, Unmarshal failed",
					logx.Field("err", err), logx.Field("payload", msg.Payload))
			}
			reminder := &models.Reminder{}
			if payload.Op != "remove" {
				err = r.db.First(reminder, payload.ID).Error
				if err != nil {
					logx.Errorw("in SubscribeChange, First failed",
						logx.Field("err", err), logx.Field("payload", msg.Payload))
					return
				}
			}
			switch payload.Op {
			case "remove":
				r.delTask(payload.ID)
			case "add":
				err = r.scheduleOneTask(reminder)
				if err != nil {
					logx.Errorw("in SubscribeChange, scheduleOneTask failed",
						logx.Field("err", err), logx.Field("payload", msg.Payload))
				}
			case "update":
				r.delTask(payload.ID)
				err = r.scheduleOneTask(reminder)
				if err != nil {
					logx.Errorw("in SubscribeChange, scheduleOneTask failed",
						logx.Field("err", err), logx.Field("payload", msg.Payload))
				}
			}

		}()

	}
}

// Publish
// {"op":"add\update\remove","id":123}
func (r *ReminderMgr) Publish(op string, id uint) {
	var payload = channelPayload{Op: op, ID: id}
	jsonData, err := json.Marshal(&payload)
	if err != nil {
		logx.Errorw("in Publish,marshal json failed", logx.Field("err", err))
		return
	}
	push, err := r.redisCmd.Publish(context.Background(),
		monitorChannelName, string(jsonData)).Result()
	if err != nil {
		logx.Errorw("in Publish, Publish failed", logx.Field("err", err),
			logx.Field("payload", string(jsonData)))
	} else {
		logx.Infow("in Publish, over,", logx.Field("push", push), logx.Field("payload", string(jsonData)))
	}
}

// func isLeap(year int) bool {
// 	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
// }

//
//
