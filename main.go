package main

import (
	logger "github.com/ipfs/go-log"
	"spike-frame/cache"
	"spike-frame/config"
	"spike-frame/dao"
	"spike-frame/global"
	"spike-frame/initialize"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	global.Viper = config.InitViper()
	global.GormClient = initialize.GormMysql()
	global.RedisClient = cache.ConnectRedis()
	global.DbAccessor = dao.NewGormAccessor(global.GormClient)
	//	chain.NewBscListener()
	initialize.RunServer()
}
