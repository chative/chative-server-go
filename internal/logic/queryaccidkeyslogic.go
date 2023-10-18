package logic

import (
	"context"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	commonlogic "chative-server-go/logic"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryAccIdKeysLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAccIdKeysLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAccIdKeysLogic {
	return &QueryAccIdKeysLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type AccountInfo struct {
	IdentityKey   string `json:"identityKey"`
	PublicConfigs struct {
		MeetingVersion int `json:"meetingVersion"`
	} `json:"publicConfigs"`
}

func (a *AccountInfo) MeetingVersion() int {
	if a.PublicConfigs.MeetingVersion == 0 {
		return 1
	}
	return a.PublicConfigs.MeetingVersion
}

type IdentityKeyInfo struct {
	UID            string `json:"uid"`
	IdentityKey    string `json:"identityKey"`
	RegistrationId uint   `json:"registrationId"`
	MsgEncVersion  int    `json:"msgEncVersion"`
}

func (l *QueryAccIdKeysLogic) QueryAccIdKeys(req *types.QueryAccIdKeysReq) (resp *types.QueryAccIdKeysResp, errInfo *response.ErrInfo) {
	redisCmd := l.svcCtx.Redis
	resp = &types.QueryAccIdKeysResp{}
	keys := make([]IdentityKeyInfo, 0, len(req.UIDs))
	meetingVersion := 0xFFFFFFFF
	for _, v := range req.UIDs {
		userInfo, err := commonlogic.GetUserBasicInfo(redisCmd, l.svcCtx.DbEngine, v)
		// data, err := redisCmd.Get(l.ctx, "Account5"+v).Result()
		if err != nil {
			l.Errorw("GetUserBasicInfo", logx.Field("err", err), logx.Field("uid", v))
			continue
		}
		keyInfo := IdentityKeyInfo{UID: v, IdentityKey: userInfo.IdentityKey}
		for _, device := range userInfo.Devices {
			if device.ID == 1 {
				keyInfo.RegistrationId = device.RegistrationId
				break
			}
		}
		keyInfo.MsgEncVersion = userInfo.MsgEncVersion()

		keys = append(keys, keyInfo)
		mv := userInfo.MeetingVersion()
		if mv < meetingVersion {
			meetingVersion = mv
		}
	}
	// arr, err := redisCmd.MGet(l.ctx, req.UIDs...).Result()
	// if err != nil {
	// 	l.Errorw("redisCmd.MGet", logx.Field("err", err), logx.Field("req", req))
	// 	errInfo = &response.ErrInfo{
	// 		ErrCode: 19999,
	// 		Reason:  "Internal error",
	// 	}
	// 	return
	// }
	// for _, v := range arr {
	// 	var accInfo AccountInfo
	// 	err = json.Unmarshal([]byte(v.(string)), &accInfo)
	// 	if err != nil {
	// 		l.Errorw("json.Unmarshal", logx.Field("err", err), logx.Field("data", v))
	// 		continue
	// 	}
	// 	keys = append(keys, IdentityKeyInfo{UID: req.UID, IdentityKey: accInfo.IdentityKey})
	// }
	resp.Keys = keys
	if len(keys) > 0 {
		resp.MeetingVersion = meetingVersion
	}

	return
}
