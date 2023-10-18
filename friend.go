package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"chative-server-go/cron"
	"chative-server-go/dbengine"
	"chative-server-go/internal/config"
	"chative-server-go/internal/handler"
	"chative-server-go/internal/svc"
	"chative-server-go/mainrpc"
	"chative-server-go/rediscluster"
	"chative-server-go/response"
	"chative-server-go/rpcserver"
	"chative-server-go/utils/iplocation"
	"chative-server-go/utils/secretsmanager"
	"chative-server-go/utils/sms"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/friend-api.yaml", "the config file")

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	var c config.Config
	conf.MustLoad(*configFile, &c)
	config.SetConfigFile(*configFile)
	if c.ReqAddFriend.AskTimeout <= 0 {
		c.ReqAddFriend.AskTimeout = 3
	}

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/utils/location" {
				next(w, r)
				return
			}
			auth := r.Header.Get("Authorization")
			if strings.HasPrefix(auth, "Basic") {
				var body = &response.Body{Ver: 1}
				body.Status = 5
				body.Reason = "Invalid token"
				w.Header().Set("errorcode", strconv.Itoa(body.Status))
				w.Header().Set("errormsg", body.Reason)
				httpx.WriteJson(w, http.StatusUnauthorized, body)
			} else {
				next(w, r)
			}
		}
	})

	iplocation.Init(c.IpLocation.Key, c.IpLocation.CountryFile)
	rediscluster.Init(c.ClusterRedis)
	sms.Init(c.SMS)
	secretsmanager.Init(c.SecretsManager)
	mainrpc.Init(c)
	ctx := svc.NewServiceContext(c)
	db, _ := dbengine.GetDB()
	cron.Init(rediscluster.GetRedis(), db)
	handler.RegisterHandlers(server, ctx)

	go rpcserver.Rpcserver(c.RpcSrvCnf)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
