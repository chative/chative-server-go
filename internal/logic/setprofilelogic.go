package logic

import (
	"context"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetProfileLogic {
	return &SetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetProfileLogic) SetProfile(req *types.SetProfileReq) (resp *types.ProfileResp, errInfo *response.ErrInfo) {
	var db = l.svcCtx.DbEngine
	var profile = models.UserProfile{UID: req.UID}
	resp = &types.ProfileResp{SearchByEmail: req.SearchByEmail,
		SearchByPhone: req.SearchByPhone, PasskeysSwitch: req.PasskeysSwitch}
	if req.PasskeysSwitch >= 0 { // 关闭passkeys
		if req.PasskeysSwitch != 0 { // 设置时只能是0
			return
		}
		err := DelWebauthnCredentials(db, l.svcCtx.Config.Webauthn.Host, req.UID)
		if err != nil {
			l.Errorw("SetProfile DelWebauthnCredentials failed", logx.Field("req", req), logx.Field("err", err))
		}
		return
	}
	err := db.First(&profile, profile).Error
	if err != nil {
		l.Errorw("SetProfile failed,find profile failed", logx.Field("uid", req.UID), logx.Field("err", err))
	}
	profile.SearchByEmail = req.SearchByEmail
	profile.SearchByPhone = req.SearchByPhone

	//
	var res *gorm.DB
	tx := db.Model(profile)
	var fields []interface{} = []interface{}{}
	if profile.SearchByEmail >= 0 {
		fields = append(fields, "search_by_email")
	}
	if profile.SearchByPhone >= 0 {
		fields = append(fields, "search_by_phone")
	}
	varargs := []interface{}{}
	if len(fields) > 1 {
		varargs = fields[1:]
	}

	res = tx.Model(profile).Select(fields[0], varargs...).
		Where(models.UserProfile{UID: profile.UID}).Updates(profile)
	if res.Error != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "internal error",
		}
		l.Errorw("SetProfile update failed", logx.Field("req", req), logx.Field("err", res.Error))
		return
	}
	if res.RowsAffected == 0 {
		// 2. 插入
		res = db.Create(&profile)
		if res.Error != nil {
			errInfo = &response.ErrInfo{
				ErrCode: 99,
				Reason:  "internal error",
			}
			l.Errorw("SetProfile insert failed", logx.Field("req", req), logx.Field("err", res.Error))
			return
		}
	}

	var hKey = logic.DirectoryHashKey
	// 更新redis
	redisCmd := l.svcCtx.Redis
	_, err = redisCmd.Pipelined(l.ctx, func(p redis.Pipeliner) error {
		if profile.EmailHash != "" {
			if req.SearchByEmail > 0 {
				p.HSet(l.ctx, hKey, profile.EmailHash, profile.UID)
			} else if req.SearchByEmail == 0 {
				p.HDel(l.ctx, hKey, profile.EmailHash)
			}
		}
		if profile.PhoneHash != "" {
			if req.SearchByPhone > 0 {
				p.HSet(l.ctx, hKey, profile.PhoneHash, profile.UID)
			} else if req.SearchByPhone == 0 {
				p.HDel(l.ctx, hKey, profile.PhoneHash)
			}
		}
		return nil
	})
	if err != nil {
		l.Errorw("SetProfile redis failed", logx.Field("req", req), logx.Field("err", err))
	}
	return
}
