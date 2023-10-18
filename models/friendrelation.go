package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type FriendRelation struct {
	ID        uint `gorm:"primarykey"`
	UpdatedAt time.Time
	UserID1   string `gorm:"column:user_id1;uniqueIndex:friend_user_id1_user_id2_uk"`
	UserID2   string `gorm:"column:user_id2;uniqueIndex:friend_user_id1_user_id2_uk"` // UserID1 < UserID2
}

func IsFriend(db *gorm.DB, userID1, userID2 string) (bool, error) {
	var friendRelation = &FriendRelation{UserID1: userID1, UserID2: userID2}
	if friendRelation.UserID1 > friendRelation.UserID2 {
		friendRelation.UserID1, friendRelation.UserID2 = friendRelation.UserID2, friendRelation.UserID1
	}
	err := db.First(friendRelation, friendRelation).Error
	if err == nil {
		return true, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}

func CreateFriendRelation(db *gorm.DB, userID1, userID2 string) error {
	var friendRelation = &FriendRelation{UserID1: userID1, UserID2: userID2}
	if friendRelation.UserID1 > friendRelation.UserID2 {
		friendRelation.UserID1, friendRelation.UserID2 = friendRelation.UserID2, friendRelation.UserID1
	}
	err := db.Create(friendRelation).Error
	return err
}
