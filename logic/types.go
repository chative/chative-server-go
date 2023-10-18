package logic

const (
	DirectoryHashKey = "directorysearch"
)

type AddReq struct {
	// InviteCode       string
	Inviter, Invitee string
}

type ExistReq struct {
	UserID1, UserID2 string
}

type ListReq struct {
	UserID string
}

type ContactNotify struct {
	NotifyType int         `json:"notifyType"`
	NotifyTime int64       `json:"notifyTime"`
	Data       ContactData `json:"data"`
	Display    int         `json:"display"`
}

type ContactData struct {
	Ver              int   `json:"ver"`
	DirectoryVersion int64 `json:"directoryVersion"`
	Members          []struct {
		Number string `json:"number"`
		Action int    `json:"action"`
	} `json:"members"`
}
