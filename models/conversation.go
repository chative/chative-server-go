package models

import (
	"time"

	"gorm.io/gorm"
)

type ShareConversationCnf struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Conversation string `gorm:"column:conversation;uniqueIndex"`

	LastOperator    string `gorm:"column:last_operator"`
	LastOperatorDid int    `gorm:"column:last_operator_did"`

	MessageExpiry int64 `gorm:"column:message_expiry;default:-1"`
	Version       int   `gorm:"column:version;default:1"`
}

type Conversation struct {
	Remark           string `gorm:"column:remark"`
	UID              string `gorm:"column:number"`
	Conversation     string `gorm:"column:conversation"`
	BlockStatus      string `gorm:"column:block_status"`
	ConfidentialMode int    `gorm:"column:confidential_mode"`
}

func (Conversation) TableName() string {
	return "conversations"
}

func ExistConversation(db *gorm.DB, uid, conversation, blockStatus string) (bool, error) {
	var count int64
	err := db.Model(&Conversation{}).Where(
		Conversation{UID: uid, Conversation: conversation, BlockStatus: blockStatus}).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
