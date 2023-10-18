package cron

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func Init(redisCmd *redis.ClusterClient, db *gorm.DB) {
	initReminder(redisCmd, db)
}
