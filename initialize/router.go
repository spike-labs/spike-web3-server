package initialize

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	v1 "spike-frame/api/v1"
	"spike-frame/middleware"
)

var log = logger.Logger("initialize")

func initRouter() *gin.Engine {
	var r = gin.Default()
	r.Use(middleware.Cors())
	publicGroup := r.Group("")
	{
		// health
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	routerGroupApp := v1.NewRouterGroup()
	queryGroup := routerGroupApp.QueryGroup
	queryApiGroup := r.Group("/query-api/v1")
	queryGroup.InitQueryGroup(queryApiGroup)

	txGroup := routerGroupApp.TxGroup
	txApiGroup := r.Group("/tx-api/v1")
	txGroup.InitTxGroup(txApiGroup)
	return r
}
