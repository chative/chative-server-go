package main

import (
	"flag"
	"log"

	"chative-server-go/dbengine"
	"chative-server-go/internal/config"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/friend-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 链接数据库
	db, err := dbengine.GetDbEngine(c)
	if err != nil {
		log.Fatal(err)
	}
	var countWeek int64
	err = db.Table("client_versions").Where("last_login > now() - INTERVAL '7 DAY'").Distinct("number").Count(&countWeek).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Println("最近7天活跃账号数:", countWeek)
	var countMonth int64
	err = db.Table("client_versions").Where("last_login > now() - INTERVAL '30 DAY'").Distinct("number").Count(&countMonth).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Println("最近30天活跃账号数:", countMonth)
	var countAll int64
	err = db.Table("client_versions").Distinct("number").Count(&countAll).Error
	if err != nil {
		log.Fatal(err)
	}
	log.Println("所有账号数:", countAll)
}
