package rediscluster

import (
	"github.com/go-redis/redis/v8"
)

var client *redis.ClusterClient

type Config struct {
	Addrs []string
}

func Init(c Config) {
	client = redis.NewClusterClient(&redis.ClusterOptions{Addrs: c.Addrs})
}

func GetRedis() *redis.ClusterClient {
	return client
}
