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

// @Summary mint nft
// @Produce json
// @Param   order_id  formData string true "game orderId"
// @Param   token_uri formData string true "nft tokenUri"
// @Param   cb        formData string true "game callBack url address"
// @Success 200       {object} response.Response
// @Failure 500       {object} response.Response
// @Router  /tx-api/v1/hotWallet/mint [post]
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

// @Summary withdraw nft
// @Produce json
// @Param   order_id         formData string true "game orderId"
// @Param   to_address       formData string true "tx toAddress"
// @Param   token_id         formData int    true "nft token id"
// @Param   contract_address formData string true "nft contract address"
// @Param   cb               formData string true "game callBack url address"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /tx-api/v1/hotWallet/withdrawNFT [post]
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

// @Summary withdraw token
// @Produce json
// @Param   order_id         formData string true "game orderId"
// @Param   to_address       formData string true "tx toAddress"
// @Param   amount           formData string true "tx token amount"
// @Param   contract_address formData string true "token contract address(native : 0x0000000000000000000000000000000000000000)"
// @Param   cb               formData string true "game callBack url address"
// @Success 200              {object} response.Response
// @Failure 500              {object} response.Response
// @Router  /tx-api/v1/hotWallet/withdrawToken [post]
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

// @Summary recharge token
// @Produce json
// @Param   order_id     formData string true "game orderId"
// @Param   from_address formData string true "tx fromAddress"
// @Param   amount       formData string true "tx token amount"
// @Param   contract_address formData string true "token contract address(native : 0x0000000000000000000000000000000000000000)"
// @Param   tx_hash      formData string true "tx hash"
// @Param   cb           formData string true "game callBack url address"
// @Success 200          {object} response.Response
// @Failure 500          {object} response.Response
// @Router  /tx-api/v1/client/rechargeToken [post]
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

// @Summary import nft
// @Produce json
// @Param   order_id     formData string true "game orderId"
// @Param   from_address formData string true "tx fromAddress"
// @Param   contract_address formData string true "nft contract address"
// @Param   token_id     formData int    true "nft token id"
// @Param   tx_hash      formData string true "tx hash"
// @Param   cb           formData string true "game callBack url address"
// @Success 200          {object} response.Response
// @Failure 500          {object} response.Response
// @Router  /tx-api/v1/client/importNft [post]
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
