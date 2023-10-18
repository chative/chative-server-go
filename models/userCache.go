package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type UserBasicInfo struct {
	PublicConfigs struct {
		MeetingVersion int    `json:"meetingVersion"`
		PublicName     string `json:"publicName"`
		MsgEncVersion  int    `json:"msgEncVersion"`
	} `json:"publicConfigs"`
	// PublicConfigs json.RawMessage `json:"publicConfigs"`
	Avatar2     string `json:"avatar2"`
	PlainName   string `json:"plainName"`
	IdentityKey string `json:"identityKey"`
	Devices     []struct {
		ID             int  `json:"id"`
		RegistrationId uint `json:"registrationId"`
	} `json:"devices"`
}

func (u *UserBasicInfo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	err := json.Unmarshal(bytes, &u)
	return err
}

// Value return json value, implement driver.Valuer interface
func (u UserBasicInfo) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u *UserBasicInfo) MeetingVersion() int {
	if u.PublicConfigs.MeetingVersion == 0 {
		return 1
	}
	return u.PublicConfigs.MeetingVersion
}

func (u *UserBasicInfo) MsgEncVersion() int {
	if u.PublicConfigs.MsgEncVersion == 0 {
		return 1
	}
	return u.PublicConfigs.MsgEncVersion
}
