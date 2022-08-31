package queryApi

import (
	"context"
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"spike-frame/request"
	"spike-frame/response"
	"spike-frame/service/queryService"
	"time"
)

var log = logger.Logger("queryApi")

const queryNftListTimeout = 20 * time.Second

type QueryGroup struct {
	manager *queryService.QueryManager
}

func NewQueryApiGroup() QueryGroup {
	return QueryGroup{
		manager: queryService.NewQueryManager(),
	}
}

func (api *QueryGroup) InitQueryGroup(g *gin.RouterGroup) {
	g.Use()
	price := g.Group("price")
	{
		price.POST("")
	}

	nft := g.Group("/nft")
	{
		nft.POST("/list", api.QueryNftList)
	}

	txRecord := g.Group("txRecord")
	{
		txRecord.POST("")
	}
}

func (api *QueryGroup) QueryNftList(c *gin.Context) {
	var service request.NftListService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s nft list", service.WalletAddress)
			ctx := context.Background()
			cctx, cancel := context.WithTimeout(ctx, queryNftListTimeout)
			result, err := api.manager.QueryNftList(cctx, service.WalletAddress)
			cancel()
			if err != nil {
				response.FailWithMessage(err.Error(), c)
				return
			}
			response.OkWithData(result, c)
		}
	} else {
		response.FailWithMessage("request params error", c)
	}
}
