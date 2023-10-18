package config

import (
	"chative-server-go/rediscluster"
	"chative-server-go/utils/secretsmanager"
	"chative-server-go/utils/sms"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

var configFile = "etc/friend-api.yaml"

func SetConfigFile(file string) {
	configFile = file
}

func LoadConfig() (c Config, err error) {
	err = conf.Load(configFile, &c)
	return
}

type Config struct {
	RpcSrvCnf zrpc.RpcServerConf
	rest.RestConf
	Dsn      string
	SMS      sms.Config
	MainGrpc struct {
		Addr      string
		PID       string
		AppID     string
		AppSecret string
	}
	ClusterRedis rediscluster.Config
	IpLocation   struct {
		Key         string
		CountryFile string
	}
	// 可选的消息过期时间配置
	MsgExpiryOpts []int64

	// 请求加好友配置
	ReqAddFriend struct {
		// 加好友请求过期时间,单位分钟
		AskTimeout int64 `json:",default=10080"`
	}

	SecretsManager secretsmanager.Config

	Upgrade struct {
		Android struct {
			AvailableVersion string `json:",default=1.0.0"`
			DownloadUrl      string
			Version          string `json:",default=1.0.0"`
			Notes            string
		}
	}
	WebAvatar struct {
		CipherHost string
		// S3Region   string
		// S3Bucket   string
	}
	WebGroupInfo struct {
		OssBucket       string `json:",optional"`
		AwsBucket       string `json:",optional"`
		AccessKeyId     string `json:",optional"`
		AccessKeySecret string `json:",optional"`
	}

	Webauthn struct {
		Host string
	}
}
