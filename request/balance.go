package request

import (
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/service/query"
)

type BalanceService struct {
	WalletAddress string `form:"wallet_address" json:"wallet_address" binding:"required"`
}

func (b *BalanceService) QueryWalletService(manager *query.BalanceService) ([]response.BalanceShow, error) {
	result, err := manager.QueryWalletService(b.WalletAddress)
	return result, err
}
