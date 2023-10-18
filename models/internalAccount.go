package models

type InternalAccount struct {
	Number     string  `gorm:"column:number;primary_key" json:"number"`
	JoinedAt   int64   `gorm:"column:joined_at;default:extract(epoch from now()) * 1000" json:"joined_at"`
	InviteCode *string `gorm:"uniqueIndex;default:null"` // 邀请码，长期有效，与uid一一对应

	// Phone  string `gorm:"column:phone" json:"phone"` // 手机号
	// Email  string `gorm:"column:email" json:"email"` // 邮箱
	Name            string `gorm:"column:name" json:"name"`
	Registered      bool   `gorm:"column:registered" json:"registered"`
	Deleted         bool   `gorm:"column:deleted" json:"deleted"`
	SupportTransfer bool   `gorm:"column:support_transfer"`
	// InviteRule int    // 0: no invite, 1: reg, 2: friend, 3: both
}

func (InternalAccount) TableName() string {
	return "internal_accounts"
}

type Account struct {
	Number string        `gorm:"column:number"`
	Data   UserBasicInfo `gorm:"column:data;type:json"`
}

func (Account) TableName() string {
	return "accounts"
}
