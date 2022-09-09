package main

import (
	logger "github.com/ipfs/go-log"
	"spike-frame/cache"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/dao"
	"spike-frame/initialize"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	constant.Viper = config.InitViper()
	constant.GormClient = initialize.GormMysql()
	constant.RedisClient = cache.ConnectRedis()
	constant.DbAccessor = dao.NewGormAccessor(constant.GormClient)
	//chain.NewBscListener()
	initialize.RunServer()
}
