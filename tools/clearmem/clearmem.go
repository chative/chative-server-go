package main

import (
	"context"
	"flag"
	"log"

	"chative-server-go/internal/config"
	"chative-server-go/rediscluster"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/friend-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	rediscluster.Init(c.ClusterRedis)

	for _, v := range c.ClusterRedis.Addrs {
		var delCnt int64 = 0
		cli := redis.NewClient(&redis.Options{
			Addr:     v,
			Password: "",
			PoolSize: 200,
		})
		keys, err := cli.Keys(context.Background(), "Account*").Result()
		if err != nil {
			log.Println("get keys failed", err, v)
			continue
		}
		for _, k := range keys {
			c, err := rediscluster.GetRedis().Del(context.Background(), k).Result()
			if err != nil {
				log.Println("del keys failed:", err, v, ",keys count:", len(keys))
			} else {
				delCnt += c
			}
		}
		log.Println("del keys success:", v, ",keys count:", delCnt)
	}
}
