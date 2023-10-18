package models

import "time"

type UserProfile struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UID       string `gorm:"uniqueIndex"`
	PhoneHash string
	EmailHash string

	SearchByEmail int
	SearchByPhone int
	// SearchFlag int `gorm:"default:0"` // 0x1: phone, 0x2: email, 0x4: xxxxxx
	// WebauthnCnt int `gorm:"default:0"` // 大于0表示已经注册过webauthn
}
