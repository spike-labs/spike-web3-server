package queryApi

import (
	"context"
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	_ "spike-frame/docs"
	"spike-frame/request"
	"spike-frame/response"
	"spike-frame/service/queryService"
	"time"
)

var log = logger.Logger("queryApi")

const queryNftListTimeout = 20 * time.Second
const queryTxRecordTimeout = 20 * time.Second

type QueryGroup struct {
	manager        *queryService.QueryManager
	balanceService *queryService.BalanceService
}

func NewQueryApiGroup() QueryGroup {
	return QueryGroup{
		manager:        queryService.NewQueryManager(),
		balanceService: queryService.BalanceSrv,
	}
}

func (api *QueryGroup) InitQueryGroup(g *gin.RouterGroup) {
	g.Use()
	{
		g.POST("/balance", api.QueryBalance)
	}

	nft := g.Group("/nft")
	{
		nft.POST("/list", api.QueryNftList)
		nft.POST("/type", api.QueryNftType)
	}

	txRecord := g.Group("txRecord")
	{
		txRecord.POST("/native", api.QueryNativeTxRecord)
		txRecord.POST("/erc20", api.QueryERC20TxRecord)
	}
}

// @Summary query single nft list
// @Produce json
// @Param   wallet_address formData string true "wallet bsc address"
// @Param   type           formData string true "nft type"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /query-api/v1/nft/list [post]
func (api *QueryGroup) QueryNftList(c *gin.Context) {
	var service request.NftListService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" || service.Type == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s nft list", service.WalletAddress)
			ctx := context.Background()
			cctx, cancel := context.WithTimeout(ctx, queryNftListTimeout)
			result, err := api.manager.QueryNftList(cctx, service.WalletAddress, service.Type)
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

// @Summary query all nft type
// @Produce json
// @Param   wallet_address formData string true "wallet bsc address"
// @Success 200            {object} response.Response
// @Failure 500            {object} response.Response
// @Router  /query-api/v1/nft/type [post]
func (api *QueryGroup) QueryNftType(c *gin.Context) {
	var service request.NftTypeService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s nft type", service.WalletAddress)
			ctx := context.Background()
			cctx, cancel := context.WithTimeout(ctx, queryNftListTimeout)
			result, err := api.manager.QueryNftType(cctx, service.WalletAddress)
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

// @Summary query wallet balance
// @Produce json
// @Param   wallet_address formData string true "wallet bsc address"
// @Success 200            {object} response.Response
// @Failure 500            {object} response.Response
// @Router  /query-api/v1/balance [post]
func (api *QueryGroup) QueryBalance(c *gin.Context) {
	var service request.BalanceService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s balance", service.WalletAddress)
			result, err := api.balanceService.QueryWalletService(service.WalletAddress)
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

// @Summary query wallet native tx list(7 days)
// @Produce json
// @Param   wallet_address formData string true "wallet bsc address"
// @Success 200            {object} response.Response
// @Failure 500            {object} response.Response
// @Router  /query-api/v1/txRecord/native [post]
func (api *QueryGroup) QueryNativeTxRecord(c *gin.Context) {
	var service request.NativeTxRecordService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s native tx record", service.WalletAddress)
			ctx := context.Background()
			cctx, cancel := context.WithTimeout(ctx, queryTxRecordTimeout)
			result, err := api.manager.QueryNativeRecord(cctx, service.WalletAddress)
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

// @Summary query wallet ERC20 tx list(7 days)
// @Produce json
// @Param   wallet_address   formData string true "wallet bsc address"
// @Param   contract_address formData string true "erc20 contract address"
// @Success 200            {object} response.Response
// @Failure 500            {object} response.Response
// @Router  /query-api/v1/txRecord/erc20 [post]
func (api *QueryGroup) QueryERC20TxRecord(c *gin.Context) {
	var service request.ERC20TxRecordService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" || service.ContractAddress == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s erc20 : %s tx record", service.WalletAddress, service.ContractAddress)
			ctx := context.Background()
			cctx, cancel := context.WithTimeout(ctx, queryTxRecordTimeout)
			result, err := api.manager.QueryERC20TxRecord(cctx, service.WalletAddress, service.ContractAddress)
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
