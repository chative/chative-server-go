package models

import "time"

const (
	ReminderRepeatNo       = 0
	ReminderRepeatDaily    = 1
	ReminderRepeatWeekly   = 2
	ReminderRepeatBiweekly = 3
	ReminderRepeatMonthly  = 4
	ReminderWeekdays       = 5
)

type Reminder struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int `gorm:"default:1"`

	Creator string

	Conversation string `json:"conversation" gorm:"index"`
	Type         string `json:"type"`
	Timezone     string `json:"timezone"`
	Timestamp    int64  `json:"timestamp"`
	NextRun      int64
	Repeat       int    `json:"repeat"`
	Description  string `json:"description"`

	// LastDayofMonth int `json:"last_day_of_month"`
}
