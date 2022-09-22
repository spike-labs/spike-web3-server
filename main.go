package main

import (
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/cache"
	"github.com/spike-engine/spike-web3-server/chain"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/global"
	"github.com/spike-engine/spike-web3-server/initialize"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	global.Viper = config.InitViper()
	global.GormClient = initialize.GormMysql()
	global.RedisClient = cache.ConnectRedis()
	global.DbAccessor = dao.NewGormAccessor(global.GormClient)
	chain.NewBscListener()
	initialize.RunServer()
}
