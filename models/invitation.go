package models

type Invitation struct {
	Code    string `gorm:"column:code" json:"code"`
	Inviter string `gorm:"column:inviter" json:"inviter"`
	Account string `gorm:"column:account" json:"account"`
	Phone   string `gorm:"column:phone" json:"phone"`
	// PhoneHash    string
	// EmailHash    string
	Timestamp    int64 `gorm:"column:timestamp" json:"timestamp"`
	RegisterTime int64 `gorm:"column:register_time" json:"register_time"`
}

func (Invitation) TableName() string {
	return "internal_accounts_invitation"
}

type Team struct {
	Name   string
	Status bool
}

func (Team) TableName() string {
	return "teams"
}
