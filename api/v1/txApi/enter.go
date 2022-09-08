package txApi

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"spike-frame/request"
	"spike-frame/response"
	"spike-frame/service/signService"
	"spike-frame/service/txService"
)

var log = logger.Logger("txApi")

type TxGroup struct {
	hwManager *signService.HotWalletManager
	txSrv     *txService.TxService
}

func NewTxGroup() (TxGroup, error) {
	hwManager, err := signService.NewHWManager()
	if err != nil {
		log.Error("===Spike log:", err)
		return TxGroup{}, err
	}
	return TxGroup{
		hwManager: hwManager,
		txSrv:     txService.TxSrv,
	}, nil
}

func (txGroup *TxGroup) InitTxGroup(g *gin.RouterGroup) {
	g.Use()
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
	err := c.ShouldBind(&service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage("request params error", c)
	}
	err = txGroup.hwManager.BatchMint(service)
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
	}

	err = txGroup.hwManager.WithdrawNFT(service)
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
	}

	err = txGroup.hwManager.WithdrawToken(service)
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

	err = txGroup.txSrv.RechargeToken(service)
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
	}

	err = txGroup.txSrv.ImportNft(service)
	if err != nil {
		log.Error("=== Spike log: ", err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
