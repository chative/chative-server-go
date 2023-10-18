package logic

import (
	"context"
	"encoding/json"

	"chative-server-go/models"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func GetUserBasicInfo(redisCmd redis.Cmdable, db *gorm.DB, uid string) (
	userInfo models.UserBasicInfo, err error) {
	cacheData, err := redisCmd.Get(context.Background(), "Account5"+uid).Result()
	if err != nil {
		var acc models.Account
		err = db.Where(models.Account{Number: uid}).First(&acc).Error
		userInfo = acc.Data
		return
	}

	err = json.Unmarshal([]byte(cacheData), &userInfo)
	if err != nil {
		return
	}
	return
}
