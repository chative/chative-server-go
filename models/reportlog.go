package models

import "time"

type ReportLog struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Informer  string `gorm:"column:informer;uniqueIndex:report_informer_suspect_uk"` // 举报人
	Suspect   string `gorm:"column:suspect;uniqueIndex:report_informer_suspect_uk"`  // 被举报人
	Type      int    `gorm:"column:type"`
	Reason    string `gorm:"column:reason"`
	Block     int    `gorm:"column:block"`
}
