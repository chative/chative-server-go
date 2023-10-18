package models

type ApnInfo struct {
	Aps struct {
		Badge string `json:"badge"`
		Alert struct {
			LocKey  string        `json:"loc-key"`
			LocArgs []interface{} `json:"loc-args"`
			Title   string        `json:"title"`
			Body    string        `json:"body"`
		} `json:"alert"`
		Sound struct {
			Volume   int    `json:"volume"`
			Critical string `json:"critical"`
		} `json:"sound"`
		Passthrough    string `json:"passthrough"`
		MutableContent int    `json:"mutable-content"`
		Msg            string `json:"msg"`
	} `json:"aps"`
}

func NewApnInfo() *ApnInfo {
	a := new(ApnInfo)
	a.Aps.Sound.Volume = 1
	a.Aps.Sound.Critical = "0"
	a.Aps.MutableContent = 1
	return a
}

func (a *ApnInfo) SetLocKey(key string) *ApnInfo {
	a.Aps.Alert.LocKey = key
	return a
}

func (a *ApnInfo) SetLocArgs(args []interface{}) *ApnInfo {
	a.Aps.Alert.LocArgs = args
	return a
}

func (a *ApnInfo) SetBody(body string) *ApnInfo {
	a.Aps.Alert.Body = body
	return a
}

func (a *ApnInfo) SetTitle(title string) *ApnInfo {
	a.Aps.Alert.Title = title
	return a
}

func (a *ApnInfo) SetPassthrough(passthrough string) *ApnInfo {
	a.Aps.Passthrough = passthrough
	return a
}

func (a *ApnInfo) SetMsg(msg string) *ApnInfo {
	a.Aps.Msg = msg
	return a
}
