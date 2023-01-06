package main

import (
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/cache"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/initialize"
	"github.com/spike-engine/spike-web3-server/service/query"
	"github.com/spike-engine/spike-web3-server/service/sign"
)

// @title Swagger Example API
// @version 0.0.1
// @description
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name api_key
// @BasePath /
func main() {
	logger.SetLogLevel("*", "INFO")
	config.Viper = config.InitViper()
	dao.GormClient = initialize.GormMysql()
	dao.DbAccessor = dao.NewGormAccessor(dao.GormClient)
	cache.RedisClient = cache.ConnectRedis()

	query.QurManager = query.NewQueryManager()
	sign.HwManager = sign.NewHWManager()
	//chain.NewBscListener()
	initialize.RunServer()
}
