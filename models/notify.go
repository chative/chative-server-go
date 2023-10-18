package models

import "encoding/json"

const (
	DTServerNotifyTypeDirectory   = 1 // 通讯录变更
	DTServerNotifyTypeShareConfig = 5 // 会话共享配置
	DTServerNotifyTypeAddContacts = 6 // 添加联系人请求
	DTServerNotifyTypeReminder    = 8 // 提醒事项
)

type Notify struct {
	NotifyType int   `json:"notifyType"`
	NotifyTime int64 `json:"notifyTime"`

	Data interface{} `json:"data"`
}

type SharingConversationCnfNotify struct {
	Operator         string `json:"operator"`
	OperatorDeviceID int    `json:"operatorDeviceId"`
	Conversation     string `json:"conversation"`
	Ver              int    `json:"ver"`
	ChangeType       int    `json:"changeType"`
	MessageExpiry    int64  `json:"messageExpiry"`
}

type AskFriendNotify struct {
	OperatorInfo struct {
		UID           string          `json:"operatorId"`
		Did           int             `json:"operatorDeviceId"`
		Name          string          `json:"operatorName"`
		Avatar        string          `json:"avatar,omitempty"`
		PublicConfigs json.RawMessage `json:"publicConfigs,omitempty"`
	} `json:"operatorInfo"`
	AskID            uint `json:"askID"`
	ActionType       int  `json:"actionType"`
	DirectoryVersion int  `json:"directoryVersion,omitempty"`
}

type DirectoryNotify struct {
	Ver              int               `json:"ver"`
	DirectoryVersion int64             `json:"directoryVersion"`
	Members          []DirectoryMember `json:"members"`
}

type DirectoryMember struct {
	Number string `json:"number"`
	Name   string `json:"name,omitempty"`
	Action int    `json:"action"`
	ExtID  int    `json:"extId"`

	Avatar        string          `json:"avatar,omitempty"`
	PublicConfigs json.RawMessage `json:"publicConfigs,omitempty"`
}

type ReminderNotify struct {
	Operator         string `json:"operator"`
	OperatorDeviceID int    `json:"operatorDeviceId"`
	Version          int    `json:"version"`
	ChangeType       int    `json:"changeType"`
	Creator          string `json:"creator"`
	Type             string `json:"type"`
	Conversation     string `json:"conversation"`
	ReminderID       uint   `json:"reminderId"`
	Timezone         string `json:"timezone"`
	Timestamp        int64  `json:"timestamp"`
	ModifyTime       int64  `json:"modifyTime"`
	Repeat           int    `json:"repeat"`
	Description      string `json:"description"`
}
