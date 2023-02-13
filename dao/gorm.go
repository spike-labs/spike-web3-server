package dao

import (
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/model"
	"gorm.io/gorm"
)

var (
	log        = logger.Logger("dao")
	GormClient *gorm.DB
	DbAccessor *GormAccessor
)

type GormAccessor struct {
	*gorm.DB
}

func NewGormAccessor(db *gorm.DB) *GormAccessor {
	return &GormAccessor{
		db,
	}
}

func (g *GormAccessor) SaveTxCb(tx model.SpikeTx) error {
	return g.DB.Create(&tx).Error
}

func (g *GormAccessor) RecordTxHash(uuidList []string, txHash string, txStatus int) error {
	return g.DB.Model(model.SpikeTx{}).Where("uuid IN ?", uuidList).Updates(model.SpikeTx{
		TxHash: txHash,
		Status: txStatus,
	}).Error
}

func (g *GormAccessor) QueryGameCb(txHash string, notifyStatus int) ([]model.SpikeTx, error) {
	var spikeTxs []model.SpikeTx
	if err := g.DB.Select("cb", "order_id", "tx_hash", "contract_address", "token_id").Where("tx_hash = ? and notify_status = ?", txHash, notifyStatus).Find(&spikeTxs).Error; err != nil {
		log.Errorf("query game cb err : %v", err)
		return spikeTxs, err
	}
	log.Infof("spikeTxs: %v ", spikeTxs)
	return spikeTxs, nil
}

func (g *GormAccessor) UpdateTxStatus(txHash string, txStatus int, payTime int64) error {
	return g.DB.Model(model.SpikeTx{}).Where("tx_hash = ?", txHash).Updates(model.SpikeTx{
		Status:  txStatus,
		PayTime: payTime,
	}).Error
}

func (g *GormAccessor) UpdateTxNotifyStatus(orderId string, notifyStatus int) error {
	return g.DB.Model(model.SpikeTx{}).Where("order_id = ?", orderId).Update("notify_status", notifyStatus).Error
}

func (g *GormAccessor) QueryNotNotifyTx(notNotifyStatus int) ([]model.SpikeTx, error) {
	var spikeTxs []model.SpikeTx
	if err := g.DB.Select("cb", "order_id", "status", "tx_hash", "create_time", "contract_address", "token_id").Where("notify_status = ?", notNotifyStatus).Find(&spikeTxs).Error; err != nil {
		return spikeTxs, err
	}
	return spikeTxs, nil
}

func (g *GormAccessor) QueryNftList(ownerAddr string, contractAddr string) ([]model.NftOwner, error) {
	var nftList []model.NftOwner
	if err := g.DB.Select("id", "owner_address", "contract_address", "token_id").Where("owner_address = ? and contract_address = ?", ownerAddr, contractAddr).Find(&nftList).Error; err != nil {
		return nftList, err
	}
	return nftList, nil
}

func (g *GormAccessor) AddNftOwner(no model.NftOwner) error {
	return g.DB.Create(&no).Error
}

func (g *GormAccessor) UpdateNftOwner(ownerAddr string, contractAddr string, tokenId int64, updateTime int64) error {
	return g.DB.Model(model.NftOwner{}).Where("contract_address = ? and token_id = ?", contractAddr, tokenId).Updates(model.NftOwner{
		UpdateTime:   updateTime,
		OwnerAddress: ownerAddr,
	}).Error
}

func (g *GormAccessor) QueryNftOwner(tokenId int64, contractAddr string) ([]model.NftOwner, error) {
	var nftOwner []model.NftOwner
	if err := g.DB.Select("id", "owner_address, update_time").Where("token_id = ? and contract_address = ?", tokenId, contractAddr).Find(&nftOwner).Error; err != nil {
		return nftOwner, err
	}
	return nftOwner, nil
}

func (g *GormAccessor) QueryApiKey() ([]model.ApiKey, error) {
	var apiKey []model.ApiKey
	if err := g.DB.Select("id", "api_key").Find(&apiKey).Error; err != nil {
		return apiKey, err
	}
	return apiKey, nil
}

func (g *GormAccessor) AddApiKey(apiKey model.ApiKey) error {
	return g.DB.Create(&apiKey).Error
}
