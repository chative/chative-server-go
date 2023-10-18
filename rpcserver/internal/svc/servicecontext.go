package svc

import (
	"chative-server-go/dbengine"
	"chative-server-go/rediscluster"
	"chative-server-go/rpcserver/internal/config"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DbEngine *gorm.DB
	Redis    redis.Cmdable
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := dbengine.GetDB()
	if err != nil {
		logx.Error("GetDbEngine failed", err)
	}
	return &ServiceContext{
		Config:   c,
		DbEngine: db,
		Redis:    rediscluster.GetRedis(),
	}
}
