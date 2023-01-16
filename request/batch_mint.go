package request

import (
	"github.com/spike-engine/spike-web3-server/service/sign"
)

type BatchMintNFTService struct {
	OrderId         string `form:"order_id" json:"order_id" binding:"required"`
	TokenURI        string `form:"token_uri" json:"token_uri" binding:"required"`
	Cb              string `form:"cb" json:"cb" binding:"required"`
	ContractAddress int    `form:"contract_address" json:"contract_address"`
}

func (b *BatchMintNFTService) BatchMint(hwManager *sign.HotWalletManager) error {
	err := hwManager.BatchMint(b.OrderId, b.TokenURI, b.Cb, b.ContractAddress)
	return err
}
