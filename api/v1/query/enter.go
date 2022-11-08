package queryApi

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	_ "github.com/spike-engine/spike-web3-server/docs"
	"github.com/spike-engine/spike-web3-server/middleware"
	"github.com/spike-engine/spike-web3-server/request"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/service/query"
)

var log = logger.Logger("queryApi")

type QueryGroup struct {
	manager        *query.QueryManager
	balanceService *query.BalanceService
}

func NewQueryApiGroup() QueryGroup {
	return QueryGroup{
		manager:        query.QurManager,
		balanceService: query.BalanceSrv,
	}
}

func (api *QueryGroup) InitQueryGroup(g *gin.RouterGroup) {
	g.Use(middleware.ApiKeyAuth())
	{
		g.POST("/balance", api.QueryBalance)
		g.POST("/nftList", api.NftList)
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

// @Summary query all type nft list
// @Produce json
// @Param   wallet_address   formData string true "wallet bsc address"
// @Param   contract_address formData string true "nft contract address"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /query-api/v1/nftList [post]
// @Security ApiKeyAuth
func (api *QueryGroup) NftList(c *gin.Context) {
	var service request.NftTypeService
	if err := c.ShouldBind(&service); err == nil {
		log.Infof("query wallet %s contractAddr %s nft list", service.WalletAddress, service.ContractAddress)
		result, err := service.NftList(api.manager)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage("request params error", c)
	}
}

// @Summary query single nft list
// @Produce json
// @Param   wallet_address   formData string true "wallet bsc address"
// @Param   contract_address formData string true "nft contract address"
// @Param   type             formData string true "nft type"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /query-api/v1/nft/list [post]
// @Security ApiKeyAuth
func (api *QueryGroup) QueryNftList(c *gin.Context) {
	var service request.NftListService
	if err := c.ShouldBind(&service); err == nil {
		log.Infof("query wallet %s contractAddr %s nft list", service.WalletAddress, service.ContractAddress)

		result, err := service.QueryNftList(api.manager)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(result, c)
	} else {
		response.FailWithMessage("request params error", c)
	}
}

// @Summary query all nft type
// @Produce json
// @Param   wallet_address   formData string true "wallet bsc address"
// @Param   contract_address formData string true "nft contract address"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /query-api/v1/nft/type [post]
// @Security ApiKeyAuth
func (api *QueryGroup) QueryNftType(c *gin.Context) {
	var service request.NftTypeService

	if err := c.ShouldBind(&service); err == nil {
		log.Infof("query wallet %s contractAddr %s nft type", service.WalletAddress, service.ContractAddress)

		result, err := service.QueryNftType(api.manager)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(result, c)

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
// @Security ApiKeyAuth
func (api *QueryGroup) QueryBalance(c *gin.Context) {
	var service request.BalanceService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s balance", service.WalletAddress)
			result, err := service.QueryWalletService(api.balanceService)
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
// @Security ApiKeyAuth
func (api *QueryGroup) QueryNativeTxRecord(c *gin.Context) {
	var service request.NativeTxRecordService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s native tx record", service.WalletAddress)

			result, err := service.QueryNativeRecord(api.manager)
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
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /query-api/v1/txRecord/erc20 [post]
// @Security ApiKeyAuth
func (api *QueryGroup) QueryERC20TxRecord(c *gin.Context) {
	var service request.ERC20TxRecordService

	if err := c.ShouldBind(&service); err == nil {
		if service.WalletAddress == "" || service.ContractAddress == "" {
			response.FailWithMessage("request params error", c)
		} else {
			log.Infof("query wallet %s erc20 : %s tx record", service.WalletAddress, service.ContractAddress)
			result, err := service.QueryERC20TxRecord(api.manager)
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
