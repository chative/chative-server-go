package response

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Ver    int         `json:"ver"`
	Status int         `json:"status"`
	Reason string      `json:"reason"`
	Data   interface{} `json:"data"`
}

type ErrInfo struct {
	ErrCode int `json:"err_code"`
	// HTTPCode int    `json:"http_code"`
	Reason string `json:"reason"`
}

func Response(w http.ResponseWriter, resp interface{}, err *ErrInfo) {
	var body = &Body{Ver: 1, Data: resp}
	if err != nil {
		body.Status = err.ErrCode
		body.Reason = err.Reason
		w.Header().Set("errorcode", strconv.Itoa(err.ErrCode))
		w.Header().Set("errormsg", err.Reason)
	} else {
		body.Reason = "OK"
	}
	httpx.OkJson(w, body)
}
