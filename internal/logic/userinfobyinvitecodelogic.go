package logic

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "sync"
	"time"

	"chative-server-go/internal/svc"
	"chative-server-go/internal/types"
	"chative-server-go/logic"
	"chative-server-go/models"
	"chative-server-go/response"

	"github.com/zeromicro/go-zero/core/logx"
)

// var (
// 	onceInitS3client sync.Once
// 	s3Client         *s3.Client
// )

type UserInfoByInviteCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoByInviteCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoByInviteCodeLogic {
	return &UserInfoByInviteCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoByInviteCodeLogic) UserInfoByInviteCode(req *types.UserInfoByInviteCodeReq) (resp *types.UserInfoByInviteCodeRes, errInfo *response.ErrInfo) {
	// onceInitS3client.Do(func() { l.initS3client() })
	//  inviteCode 查uid
	db := l.svcCtx.DbEngine
	if len(req.InviteCode) > 8 {
		req.InviteCode = req.InviteCode[:8]
	}
	var accountInternal = models.InternalAccount{InviteCode: &req.InviteCode}
	err := db.Where(accountInternal).First(&accountInternal).Error
	if err != nil {
		l.Errorw("UserInfoByInviteCode failed",
			logx.Field("inviteCode", req.InviteCode), logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 10107,
			Reason:  "Expired Invite Code",
		}
		return
	}
	resp = &types.UserInfoByInviteCodeRes{Name: accountInternal.Name}
	defer func() {
		if resp.AvatarContent == "" {
			resp.AvatarContent = defaultWebAvatar
		}
	}()
	// uid查头像，
	userInfo, err := logic.GetUserBasicInfo(l.svcCtx.Redis, db, accountInternal.Number)
	if err != nil {
		l.Errorw("UserInfoByInviteCode GetUserBasicInfo failed",
			logx.Field("inviteCode", req.InviteCode), logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Expired Invite Code",
		}
		return
	}
	// {"encAlgo":"AESGCM256","encKey":"1f5Tofm23dcJ9ZsHFWjNif8iZoY7BJ6HE6YVeqcTjTM=","attachmentId":"7042021153121170216"}
	var avatarInfo struct {
		AttachmentId string `json:"attachmentId"`
		EncKey       string `json:"encKey"`
	}
	if len(userInfo.Avatar2) <= 2 {
		return
	}
	err = json.Unmarshal([]byte(userInfo.Avatar2), &avatarInfo)
	if err != nil {
		l.Errorw("UserInfoByInviteCode json.Unmarshal failed",
			logx.Field("uid", accountInternal.Number),
			logx.Field("inviteCode", req.InviteCode), logx.Field("err", err))
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Expired Invite Code",
		}
		return
	}
	if avatarInfo.AttachmentId == "" {
		return
	}

	redisCmd := l.svcCtx.Redis
	redisKey := "webavatar:" + avatarInfo.AttachmentId
	avatar, err := redisCmd.Get(l.ctx, redisKey).Bytes()
	if err == nil {
		resp.AvatarContent = base64.RawStdEncoding.EncodeToString(avatar)
		return
	}
	plainData, err := l.decAvatar(avatarInfo.AttachmentId, avatarInfo.EncKey)
	if err != nil {
		errInfo = &response.ErrInfo{
			ErrCode: 99,
			Reason:  "Expired Invite Code",
		}
		return
	}
	err = redisCmd.Set(l.ctx, redisKey, plainData, time.Hour*24*7).Err()
	if err != nil {
		l.Errorw("UserInfoByInviteCode redisCmd.Set failed", logx.Field("err", err))
	} else {
		l.Infow("UserInfoByInviteCode redisCmd.Set success", logx.Field("redisKey", redisKey))
	}
	resp.AvatarContent = base64.RawStdEncoding.EncodeToString(plainData)
	return
}

// func genAvatarPath(attID string) string {
// 	return "webavatar/" + attID
// }

func (l *UserInfoByInviteCodeLogic) decAvatar(attID string, encKey string) (plainData []byte, err error) {
	//
	key, err := base64.StdEncoding.DecodeString(encKey)
	if err != nil {
		l.Errorw("decUplaodAvatar base64.StdEncoding.DecodeString failed",
			logx.Field("attID", attID), logx.Field("encKey", encKey), logx.Field("err", err))
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		l.Errorw("decUplaodAvatar aes.NewCipher failed",
			logx.Field("attID", attID), logx.Field("encKey", encKey), logx.Field("err", err))
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		l.Errorw("decUplaodAvatar aes.NewGCM failed",
			logx.Field("attID", attID), logx.Field("encKey", encKey), logx.Field("err", err))
		return
	}
	webAvatar := l.svcCtx.Config.WebAvatar
	resp, err := http.Get(webAvatar.CipherHost + "/" + attID)
	if err != nil {
		l.Errorw("decUplaodAvatar http.Get failed",
			logx.Field("attID", attID), logx.Field("encKey", encKey), logx.Field("err", err))
		return
	}
	defer resp.Body.Close()
	nonceSize := gcm.NonceSize()
	allEncData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l.Errorw("decUplaodAvatar ioutil.ReadAll failed",
			logx.Field("attID", attID), logx.Field("encKey", encKey), logx.Field("err", err))
		return
	}
	plainData, err = gcm.Open(nil, allEncData[:nonceSize], allEncData[nonceSize:], nil)
	// if err != nil {
	// 	l.Errorw("decUplaodAvatar gcm.Open failed",
	// 		logx.Field("attID", attID), logx.Field("encKey", encKey), logx.Field("err", err))
	// 	return
	// }
	// //
	// objKey := genAvatarPath(attID)
	// _, err = s3Client.PutObject(context.Background(),
	// 	&s3.PutObjectInput{Bucket: &webAvatar.S3Bucket, Key: &objKey,
	// 		Body:        bytes.NewReader(plainData),
	// 		ContentType: aws.String("image/png")}, func(o *s3.Options) {
	// 		o.UseAccelerate = true
	// 		o.RetryMaxAttempts = 3
	// 	})
	// if err != nil {
	// 	l.Errorw("decUplaodAvatar s3Client.PutObject failed",
	// 		logx.Field("attID", attID), logx.Field("encKey", encKey), logx.Field("err", err))
	// 	return
	// }
	return
}

// func (l *UserInfoByInviteCodeLogic) initS3client() {
// 	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		l.Errorw("initS3client LoadDefaultConfig failed", logx.Field("err", err))
// 		return
// 	}
// 	webavatarCfg := l.svcCtx.Config.WebAvatar
// 	sdkConfig.Region = webavatarCfg.S3Region
// 	s3Client = s3.NewFromConfig(sdkConfig)
// }
