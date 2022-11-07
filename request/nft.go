package request

import (
	"context"
	"github.com/spike-engine/spike-web3-server/model"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/service/query"
	"time"
)

const queryNftListTimeout = 20 * time.Second
const queryTxRecordTimeout = 20 * time.Second

type NftListService struct {
	WalletAddress   string `form:"wallet_address" json:"wallet_address" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
	Type            string `form:"type" json:"type" binding:"required"`
}

func (n *NftListService) QueryNftList(manager *query.QueryManager) ([]model.CacheData, error) {
	ctx := context.Background()
	cctx, cancel := context.WithTimeout(ctx, queryNftListTimeout)
	result, err := manager.QueryNftList(cctx, n.ContractAddress, n.WalletAddress, n.Type)
	cancel()
	return result, err
}

type NftTypeService struct {
	WalletAddress   string `form:"wallet_address" json:"wallet_address" binding:"required"`
	ContractAddress string `form:"contract_address" json:"contract_address" binding:"required"`
}

func (n *NftTypeService) NftList(manager *query.QueryManager) ([]model.CacheData, error) {
	result, err := manager.NftList(n.ContractAddress, n.WalletAddress)
	return result, err
}

func (n *NftTypeService) QueryNftType(manager *query.QueryManager) ([]response.NftType, error) {
	ctx := context.Background()
	cctx, cancel := context.WithTimeout(ctx, queryNftListTimeout)
	result, err := manager.QueryNftType(cctx, n.ContractAddress, n.WalletAddress)
	cancel()
	return result, err
}
