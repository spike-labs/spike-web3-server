package queryApi

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/middleware"
	"github.com/spike-engine/spike-web3-server/request"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/service/query"
	"github.com/spike-engine/spike-web3-server/util"
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

	crypto := g.Group("crypto")
	{
		crypto.POST("generateEcdsaKeyPair", api.GenerateKeyPair)
		crypto.POST("ecdsaSign", api.EcdsaSign)
	}
}

func (api *QueryGroup) GenerateKeyPair(c *gin.Context) {
	privateKey, publicKey, err := util.GenerateEcdsaKey()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(struct {
		PrivateKey string `json:"private_key"`
		PublicKey  string `json:"public_key"`
	}{privateKey, publicKey}, c)
}

func (api *QueryGroup) EcdsaSign(c *gin.Context) {
	var service request.SignRequest
	if err := c.ShouldBindJSON(&service); err == nil {
		log.Infof("%s sign msg: %s", service.PrivateKey, service.SignMsg)
		signature, err := util.EcdsaSign(service.PrivateKey, service.SignMsg)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithData(struct {
			Signature string `json:"signature"`
		}{
			signature,
		}, c)
	} else {
		response.FailWithMessage("request params error", c)
	}
}

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
