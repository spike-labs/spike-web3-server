package request

import (
	"context"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/service/query"
)

type NativeTxRecordService struct {
	WalletAddress string `form:"wallet_address" json:"wallet_address" binding:"required"`
}

func (n *NativeTxRecordService) QueryNativeRecord(manager *query.QueryManager) (response.BscResult, error) {
	ctx := context.Background()
	cctx, cancel := context.WithTimeout(ctx, queryTxRecordTimeout)
	result, err := manager.QueryNativeRecord(cctx, n.WalletAddress)
	cancel()
	return result, err
}

type ERC20TxRecordService struct {
	WalletAddress   string `form:"wallet_address" json:"wallet_address" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
}

func (e *ERC20TxRecordService) QueryERC20TxRecord(manager *query.QueryManager) (response.BscResult, error) {
	ctx := context.Background()
	cctx, cancel := context.WithTimeout(ctx, queryTxRecordTimeout)
	result, err := manager.QueryERC20TxRecord(cctx, e.WalletAddress, e.ContractAddress)
	cancel()
	return result, err
}
