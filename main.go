package main

import (
	logger "github.com/ipfs/go-log"
	"spike-frame/cache"
	"spike-frame/chain"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/initialize"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	constant.Viper = config.InitViper()
	constant.GormClient = initialize.GormMysql()
	constant.RedisClient = cache.ConnectRedis()
	chain.NewBscListener()
	initialize.RunServer()
}
