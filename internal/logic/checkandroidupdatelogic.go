package logic

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"chative-server-go/internal/config"
	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/response"

	"github.com/go-redis/redis/v8"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAndroidUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckAndroidUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAndroidUpdateLogic {
	return &CheckAndroidUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const (
	androidUpgradeInfoKey = "android_upgrade_info"
)

func (l *CheckAndroidUpdateLogic) CheckAndroidUpdate(req *types.CheckUpdateReq) (resp *types.CheckUpdateRes, errInfo *response.ErrInfo) {
	redisCmd := l.svcCtx.Redis
	var androidUpgradeInfo struct {
		AvailableVersion string `redis:"available_version"`
		Version          string `redis:"version"`
		URL              string `redis:"url"`
		Notes            string `redis:"notes"`
	}
	err := redisCmd.HGetAll(l.ctx, androidUpgradeInfoKey).Scan(&androidUpgradeInfo)
	if err != nil || androidUpgradeInfo.AvailableVersion == "" {
		if err != nil && !errors.Is(err, redis.Nil) { // 其他错误
			l.Errorw("in CheckAndroidUpdate redis HGetAll error", logx.Field("err", err))
		}
		c, err := config.LoadConfig()
		if err != nil {
			l.Errorw("in CheckAndroidUpdate LoadConfig error", logx.Field("err", err))
			errInfo = &response.ErrInfo{
				ErrCode: 19999,
				Reason:  "Config error",
			}
			return
		}
		// AvailableVersion string `json:",default=1.0.0"`
		// DownloadUrl      string
		// Version          string `json:",default=1.0.0"`
		// Notes            string
		android := c.Upgrade.Android
		androidUpgradeInfo.AvailableVersion = android.AvailableVersion
		androidUpgradeInfo.Version = android.Version
		androidUpgradeInfo.URL = android.DownloadUrl
		androidUpgradeInfo.Notes = android.Notes
		err = redisCmd.HSet(l.ctx, androidUpgradeInfoKey, map[string]interface{}{
			"available_version": android.AvailableVersion,
			"version":           android.Version,
			"url":               android.DownloadUrl,
			"notes":             android.Notes,
		}).Err()
		if err != nil {
			l.Errorw("in CheckAndroidUpdate redis HSet error", logx.Field("err", err))
		} else {
			l.Infow("in CheckAndroidUpdate redis HSet success", logx.Field("androidUpgradeInfo", androidUpgradeInfo))
		}
	}
	resp = &types.CheckUpdateRes{
		Update: false,
		Force:  false,
		Url:    androidUpgradeInfo.URL,
		Notes:  androidUpgradeInfo.Notes,
	}
	// 比对 version
	curMajor, curMinor, curPatch, err := ParseVersion(req.CurVersion)
	if err != nil {
		if errors.Is(err, ErrVersionFormat) {
			errInfo = &response.ErrInfo{
				ErrCode: 1,
				Reason:  "Invalid parameter",
			}
			return
		}
		l.Errorw("in CheckAndroidUpdate ParseVersion error", logx.Field("err", err), logx.Field("req", req))
		return
	}
	major, minor, patch, err := ParseVersion(androidUpgradeInfo.Version)
	if err != nil {
		l.Errorw("in CheckAndroidUpdate ParseVersion error", logx.Field("err", err), logx.Field("androidUpgradeInfoVersion", androidUpgradeInfo.Version))
		return
	}
	if curMajor < major ||
		(curMajor == major && curMinor < minor) ||
		(curMajor == major && curMinor == minor && curPatch < patch) {
		resp.Update = true
	} else {
		l.Infow("in CheckAndroidUpdate no update", logx.Field("req", req), logx.Field("androidUpgradeInfo", androidUpgradeInfo))
		return
	}
	// 比对 available_version
	avaMajor, avaMinor, avaPatch, err := ParseVersion(androidUpgradeInfo.AvailableVersion)
	if err != nil {
		l.Errorw("in CheckAndroidUpdate ParseVersion error", logx.Field("err", err), logx.Field("androidUpgradeInfoAvailableVersion", androidUpgradeInfo.AvailableVersion))
		return
	}
	if curMajor < avaMajor ||
		(curMajor == avaMajor && curMinor < avaMinor) ||
		(curMajor == avaMajor && curMinor == avaMinor && curPatch < avaPatch) {
		resp.Force = true
	}
	l.Infow("in CheckAndroidUpdate", logx.Field("req", req), logx.Field("androidUpgradeInfo", androidUpgradeInfo), logx.Field("resp", resp))
	return
}

var (
	ErrVersionFormat = errors.New("version format error")
)

func ParseVersion(version string) (major, minor, patch int, err error) {
	vs := strings.Split(version, ".")
	if len(vs) != 3 {
		err = ErrVersionFormat
		return
	}

	v := strings.TrimLeft(strings.TrimSpace(vs[0]), "0")
	if v != "" {
		major, err = strconv.Atoi(v)
		if err != nil {
			return
		}
	}
	v = strings.TrimLeft(strings.TrimSpace(vs[1]), "0")
	if v != "" {
		minor, err = strconv.Atoi(v)
		if err != nil {
			return
		}
	}
	v = strings.TrimLeft(strings.TrimSpace(vs[2]), "0")
	if v != "" {
		patch, err = strconv.Atoi(v)
	}
	return
}
