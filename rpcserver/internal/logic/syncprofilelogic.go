package logic

import (
	"context"

	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/rpcserver/internal/svc"
	"chative-server-go/rpcserver/pb"
	"chative-server-go/utils/crypto"
	"chative-server-go/utils/secretsmanager"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncProfileLogic {
	return &SyncProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SyncProfileLogic) SyncProfile(in *pb.SyncProfileRequest) (resp *pb.SyncProfileResponse, err error) {
	var db = l.svcCtx.DbEngine
	sm := secretsmanager.GetSM()

	respBase := pb.BaseResponse{
		Ver: 1, Status: 0, Reason: "OK",
	}
	resp = &pb.SyncProfileResponse{Base: &respBase}
	var profile = models.UserProfile{UID: in.Uid}
	db.Model(profile).Where(profile).First(&profile)
	var fields []interface{} = []interface{}{}
	if in.Email != "" {
		profile.EmailHash, err = crypto.HashID(ClearEmail(in.Email), sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("SyncProfile hash email failed", logx.Field("uid", profile.UID), logx.Field("err", err))
		}
		fields = append(fields, "email_hash")
	} else if in.EmailHash != "" {
		profile.EmailHash = in.EmailHash
		fields = append(fields, "email_hash")
	}

	if in.Phone != "" {
		profile.PhoneHash, err = crypto.HashID(in.Phone, sm.GetDirectoryClientSalt(), sm.GetDirectoryServerSalt())
		if err != nil {
			l.Errorw("SyncProfile hash phone failed", logx.Field("uid", profile.UID), logx.Field("err", err))
		}
		fields = append(fields, "phone_hash")
	} else if in.PhoneHash != "" {
		profile.PhoneHash = in.PhoneHash
		fields = append(fields, "phone_hash")
	}

	// if len(fields) == 0 {
	// 	l.Errorw("SyncProfile no fields to update", logx.Field("uid", profile.UID))
	// 	return
	// }

	var res *gorm.DB
	varargs := []interface{}{}
	if len(fields) > 1 {
		varargs = fields[1:]
	}

	// 1. 需要更新
	if len(fields) > 0 {
		tx := db.Model(profile)
		res = tx.Select(fields[0], varargs...).Where(models.UserProfile{UID: profile.UID}).
			Updates(profile)
		if res.Error != nil {
			respBase.Status = 19999
			respBase.Reason = res.Error.Error()
			l.Errorw("SyncProfile update failed", logx.Field("uid", profile.UID), logx.Field("err", res.Error))
			return
		}
	}

	// 2. 没有对应的记录就插入
	if profile.ID == 0 {
		// 2. 插入
		res = db.Create(&profile)
		if res.Error != nil {
			respBase.Status = 19999
			respBase.Reason = res.Error.Error()
			l.Errorw("SyncProfile insert failed", logx.Field("uid", profile.UID), logx.Field("err", res.Error))
			return
		}
	}

	var hKey = logic.DirectoryHashKey
	// 更新redis
	redisCmd := l.svcCtx.Redis
	_, err = redisCmd.Pipelined(l.ctx, func(p redis.Pipeliner) error {
		if profile.EmailHash != "" {
			if profile.SearchByEmail > 0 {
				p.HSet(l.ctx, hKey, profile.EmailHash, profile.UID)
			} else if profile.SearchByEmail == 0 {
				p.HDel(l.ctx, hKey, profile.EmailHash)
			}
		}
		if profile.PhoneHash != "" {
			if profile.SearchByPhone > 0 {
				p.HSet(l.ctx, hKey, profile.PhoneHash, profile.UID)
			} else if profile.SearchByPhone == 0 {
				p.HDel(l.ctx, hKey, profile.PhoneHash)
			}
		}
		return nil
	})
	l.Infow("SyncProfile redis", logx.Field("uid", profile.UID), logx.Field("err", err),
		logx.Field("EmailHash", profile.EmailHash), logx.Field("PhoneHash", profile.PhoneHash))
	if err != nil {
		respBase.Status = 19999
		respBase.Reason = err.Error()
		l.Errorw("SyncProfile redis failed", logx.Field("uid", profile.UID), logx.Field("err", err))
	}
	return
}
