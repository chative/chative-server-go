package models

import "gorm.io/gorm"

type Group struct {
	ID     string `gorm:"column:id"`
	Name   string `gorm:"column:name"`
	Avatar string `gorm:"column:avatar"`

	Status           int  `gorm:"column:status"`          // 0: normal, other: abnormal
	InvitationRule   int  `gorm:"column:invitation_rule"` // invitation_rule
	LinkInviteSwitch bool `gorm:"column:link_invite_switch;default:true"`
}

func (Group) TableName() string {
	return "groups"
}

type GroupMember struct {
	Gid        string  `gorm:"column:gid"`
	Uid        string  `gorm:"column:uid"`
	Role       int     `gorm:"column:role"`
	InviteCode *string `gorm:"uniqueIndex;default:null"`
}

func (GroupMember) TableName() string {
	return "group_members"
}

func IsGroupMember(db *gorm.DB, gid, uid string) (ok bool, err error) {
	var count int64
	err = db.Model(&GroupMember{}).Where("gid = ? and uid = ?", gid, uid).
		Count(&count).Error
	return count > 0, err
}
