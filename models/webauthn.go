package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type WebauthnUser struct {
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        uuid.UUID `gorm:"type:uuid;primary_key"` // webauthn的user id
	Password      string
	CredentialCnt int    `gorm:"column:credential_cnt;default:0"` // 大于0表示已经注册过webauthn
	ChatUID       string `gorm:"uniqueIndex"`                     // chat uid
}
