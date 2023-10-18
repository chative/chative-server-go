package models

import (
	"time"

	"gorm.io/gorm"
)

type AskNewFriend struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Inviter    string
	InviterDid int
	Invitee    string
	Status     int `gorm:"default:1"` // 1. pending, 2. accepted, 3. rejected
}

func FirstValidAsk(db *gorm.DB, inviter, invitee string) (*AskNewFriend, error) {
	var askNewFriend = &AskNewFriend{Inviter: inviter, Invitee: invitee, Status: 1}
	err := db.Where(askNewFriend).Order("id desc").First(askNewFriend).Error
	return askNewFriend, err
}
