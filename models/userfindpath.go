package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type UserFindPath struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Src       string   `gorm:"column:src;uniqueIndex:userfindpath_src_dst_uk"`
	Dst       string   `gorm:"column:dst;uniqueIndex:userfindpath_src_dst_uk"`
	Path      FindPath `gorm:"column:path;type:json"`
}

type FindPath struct {
	Type    string `json:"type"` // fromGroup ,shareContact, link, search
	GroupID string `json:"groupID"`
	UID     string `json:"uid"` // for shareContact
}

func (u *FindPath) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, &u)
	return err
}

// Value return json value, implement driver.Valuer interface
func (u FindPath) Value() (driver.Value, error) {
	return json.Marshal(u)
}
