package txApi

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/middleware"
	"github.com/spike-engine/spike-web3-server/request"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/service/sign"
	"github.com/spike-engine/spike-web3-server/service/tx"
)

var log = logger.Logger("txApi")

type TxGroup struct {
	hwManager *sign.HotWalletManager
	txSrv     *tx.TxService
}

func NewTxGroup() (TxGroup, error) {
	return TxGroup{
		hwManager: sign.HwManager,
		txSrv:     tx.TxSrv,
	}, nil
}

func (txGroup *TxGroup) InitTxGroup(g *gin.RouterGroup) {
	g.Use(middleware.WhiteListAuth())
	hotWallet := g.Group("hotWallet")
	{
		hotWallet.POST("/mint", txGroup.BatchMint)
		hotWallet.POST("/withdrawNFT", txGroup.BatchWithdrawNFT)
		hotWallet.POST("withdrawToken", txGroup.BatchWithdrawToken)
	}
	client := g.Group("client")
	{
		client.POST("/rechargeToken", txGroup.RechargeToken)
		client.POST("/importNft", txGroup.ImportNft)
	}
}

func (txGroup *TxGroup) BatchMint(c *gin.Context) {
	var service request.BatchMintNFTService
	err := c.ShouldBindJSON(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}
	err = service.BatchMint(txGroup.hwManager)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (txGroup *TxGroup) BatchWithdrawNFT(c *gin.Context) {
	var service request.BatchWithdrawalNFTService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}

	err = service.WithdrawNFT(txGroup.hwManager)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (txGroup *TxGroup) BatchWithdrawToken(c *gin.Context) {
	var service request.BatchWithdrawalTokenService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}

	err = service.WithdrawToken(txGroup.hwManager)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (txGroup *TxGroup) RechargeToken(c *gin.Context) {
	var service request.RechargeTokenService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}

	err = service.RechargeToken(txGroup.txSrv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (txGroup *TxGroup) ImportNft(c *gin.Context) {
	var service request.ImportNftService
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
		return
	}

	err = service.ImportNft(txGroup.txSrv)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
