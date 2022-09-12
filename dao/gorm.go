package dao

import (
	logger "github.com/ipfs/go-log"
	"gorm.io/gorm"
	"spike-frame/model"
)

var log = logger.Logger("dao")

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

func (g *GormAccessor) QueryGameCb(txHash string) ([]model.SpikeTx, error) {
	var spikeTxs []model.SpikeTx
	if err := g.DB.Select("cb", "order_id").Where("tx_hash = ?", txHash).Find(&spikeTxs).Error; err != nil {
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

func (g *GormAccessor) UpdateTxNotifyStatus(orderId string, notifyStatus int64) error {
	return g.DB.Model(model.SpikeTx{}).Where("order_id = ?", orderId).Update("notify_status", notifyStatus).Error
}

func (g *GormAccessor) QueryNotNotifyTx(notNotifyStatus int) ([]model.SpikeTx, error) {
	var spikeTxs []model.SpikeTx
	if err := g.DB.Select("cb", "order_id", "status").Where("notify_status = ?", notNotifyStatus).Find(&spikeTxs).Error; err != nil {
		return spikeTxs, err
	}
	return spikeTxs, nil
}
