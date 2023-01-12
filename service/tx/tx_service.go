package tx

import (
	"time"

	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/model"
)

type TxService struct {
}

var TxSrv = new(TxService)

func (t *TxService) RechargeToken(orderId string, fromAddress string, amount string, contractAddress int, txHash string, cb string) error {
	tx := model.SpikeTx{
		OrderId:         orderId,
		From:            fromAddress,
		To:              config.Cfg.Contract.GameVaultAddress,
		Amount:          amount,
		ContractAddress: config.Cfg.Contract.ERC20ContractAddress[contractAddress],
		TxHash:          txHash,
		Cb:              cb,
		CreateTime:      time.Now().UnixMilli(),
	}
	return dao.DbAccessor.SaveTxCb(tx)
}

func (t *TxService) ImportNft(orderId string, from string, contractAddress int, tokenId int64, txHash string, cb string) error {
	tx := model.SpikeTx{
		OrderId:         orderId,
		From:            from,
		To:              config.Cfg.Contract.GameVaultAddress,
		TokenId:         tokenId,
		ContractAddress: config.Cfg.Contract.NftContractAddress[contractAddress],
		TxHash:          txHash,
		Cb:              cb,
		CreateTime:      time.Now().UnixMilli(),
	}
	return dao.DbAccessor.SaveTxCb(tx)
}
