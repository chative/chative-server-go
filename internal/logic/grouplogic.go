package logic

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"sync"
	"time"

	"chative-server-go/internal/config"
	"chative-server-go/models"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	configs3 "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

var (
	clientOSS            *oss.Client
	s3Client             *s3.Client
	ossBucket, awsBucket string
	onceInit             sync.Once
)

// 1. 查inviteCode对应的group、groupMember（name）
type InviteInfo struct {
	Group       models.Group
	GroupMember models.GroupMember
}

func InitOSSclient(cfg *config.Config) {
	onceInit.Do(func() {
		ossBucket = cfg.WebGroupInfo.OssBucket
		awsBucket = cfg.WebGroupInfo.AwsBucket
		if ossBucket == "" {
			sdkConfig, err := configs3.LoadDefaultConfig(context.TODO())
			if err != nil {
				panic(err)
				return
			}
			sdkConfig.Region = "ap-southeast-1"
			s3Client = s3.NewFromConfig(sdkConfig)
			return
		}
		var err error
		clientOSS, err = oss.New("https://oss-accelerate.aliyuncs.com",
			cfg.WebGroupInfo.AccessKeyId, cfg.WebGroupInfo.AccessKeySecret)
		if err != nil {
			panic(err)
		}
	})
}

func GetInviteInfo(redisCmd redis.Cmdable, db *gorm.DB, inviteCode string) (
	inviteInfo InviteInfo, errCode int, err error) {
	err = db.Where(models.GroupMember{InviteCode: &inviteCode}).
		First(&inviteInfo.GroupMember).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errCode = 10120
			return
		}
		errCode = 99
		return
	}
	err = db.Where(models.Group{ID: inviteInfo.GroupMember.Gid}).First(&inviteInfo.Group).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errCode = 10124
			return
		}
		errCode = 99
		return
	}
	// if inviteInfo.Group.Status != 0 {
	// 	errCode = 10123
	// 	return
	// }
	return
}

// 2. 是否启用了邀请码入群，是否有邀请权限
func CheckInvitePermission(info InviteInfo) (errCode int) {
	if !info.Group.LinkInviteSwitch {
		errCode = 10121
		return
	}
	if info.Group.InvitationRule != 2 && info.GroupMember.Role == 2 {
		errCode = 10122
		return
	}
	return
}

// 3. 有avata时，解密avatar
func FetchAvatar(avatar string, redisCmd redis.Cmdable, logger logx.Logger) (content string, err error) {
	var data []byte
	var avataData struct {
		Data string `json:"data"`
	}
	err = json.Unmarshal([]byte(avatar), &avataData)
	if err != nil || avataData.Data == "" {
		logger.Errorw("json.Unmarshal Error:", logx.Field("err", err), logx.Field("avatar", avatar))
		return "", err
	}
	var avatarMeta struct {
		ByteCount      string `json:"byteCount"`
		Digest         string `json:"digest"`
		EncryptionKey  string `json:"encryptionKey"`
		ServerID       string `json:"serverId"`
		AttachmentType int    `json:"attachmentType"`
		ContentType    string `json:"contentType"`
	}
	avatarMetaBin, err := base64.StdEncoding.DecodeString(avataData.Data)
	if err != nil {
		logger.Errorw("base64.StdEncoding.DecodeString Error:", logx.Field("err", err), logx.Field("avatar", avatar))
		return "", err
	}
	err = json.Unmarshal(avatarMetaBin, &avatarMeta)
	if err != nil {
		logger.Errorw("json.Unmarshal Error:", logx.Field("err", err), logx.Field("avatar", avatar))
		return "", err
	}
	// 查缓存
	var cacheKey = "web:group:avatar:" + avatarMeta.ServerID
	plainAvatar, err := redisCmd.Get(context.Background(), cacheKey).Bytes()
	if err == nil {
		content = base64.RawStdEncoding.EncodeToString(plainAvatar)
		return
	}

	// 下载解密
	if clientOSS != nil {
		// 获取阿里云 OSS 中的 Object
		bucketOSS, errInternal := clientOSS.Bucket(ossBucket)
		if errInternal != nil {
			logger.Errorw("clientOSS.Bucket Error:", logx.Field("err", errInternal), logx.Field("ossBucket", ossBucket))
			return "", errInternal
		}
		obj, errInternal := bucketOSS.GetObject(avatarMeta.ServerID)
		if errInternal != nil {
			logger.Errorw("bucketOSS.GetObject Error:", logx.Field("err", errInternal), logx.Field("ossBucket", ossBucket))
			return "", errInternal
		}
		defer obj.Close()
		data, err = ioutil.ReadAll(obj)
	} else {
		output, errInternal := s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String(awsBucket),
			Key:    aws.String(avatarMeta.ServerID),
		})
		if errInternal != nil {
			logger.Errorw("s3Client.GetObject Error:", logx.Field("err", errInternal), logx.Field("awsBucket", awsBucket))
			return "", errInternal
		}
		defer output.Body.Close()
		data, err = ioutil.ReadAll(output.Body)
	}
	if err != nil {
		logger.Errorw("ioutil.ReadAll Error:", logx.Field("err", err), logx.Field("avatar", avatar))
		return "", err
	}

	encKey, err := base64.StdEncoding.DecodeString(avatarMeta.EncryptionKey)
	if err != nil {
		logger.Errorw("base64.StdEncoding.DecodeString Error:", logx.Field("err", err), logx.Field("avatar", avatar))
	}
	plainAvatar, err = DecryptAESCBC(data[16:], encKey[:32], data[:16])
	if err != nil {
		logger.Errorw("DecryptAESCBC Error:", logx.Field("err", err), logx.Field("avatar", avatar))
		return "", err
	}
	content = base64.RawStdEncoding.EncodeToString(plainAvatar)
	// 写缓存
	err = redisCmd.Set(context.Background(), cacheKey, plainAvatar, time.Hour*24*7).Err()
	if err != nil {
		logger.Errorw("redisCmd.Set Error:", logx.Field("err", err), logx.Field("avatar", avatar))
	}
	return
}

// PKCS5UnPadding removes PKCS5 padding from data.
func PKCS5UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// DecryptAESCBC decrypts the given encrypted data using AES-128-CBC with PKCS5 padding.
func DecryptAESCBC(encrypted, key, iv []byte) ([]byte, error) {
	// 创建加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 检查IV长度是否正确
	if len(iv) != block.BlockSize() {
		return nil, err
	}

	// 创建CBC解密器
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密数据
	mode.CryptBlocks(encrypted, encrypted)

	// 移除填充
	plaintext := PKCS5UnPadding(encrypted)
	return plaintext, nil
}
