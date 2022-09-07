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

func (g *GormAccessor) SaveTxCb(orderId, uuid, txHash, from, to, amount string, tokenId int64, cb string, txType, createTime int64) error {
	tx := model.SpikeTx{
		OrderId:    orderId,
		Uuid:       uuid,
		TxHash:     txHash,
		From:       from,
		To:         to,
		Amount:     amount,
		TokenId:    tokenId,
		Cb:         cb,
		TxType:     txType,
		CreateTime: createTime,
	}
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
	if err := g.DB.Select("cb").Where("tx_hash = ?", txHash).Find(&spikeTxs).Error; err != nil {
		log.Errorf("query game cb err : %v", err)
		return spikeTxs, err
	}
	return spikeTxs, nil
}

func (g *GormAccessor) UpdateTxStatus(txHash string, txStatus int, payTime int64) error {
	return g.DB.Model(model.SpikeTx{}).Where("tx_hash = ?", txHash).Updates(model.SpikeTx{
		Status:  txStatus,
		PayTime: payTime,
	}).Error
}

func (g *GormAccessor) UpdateWithdrawTxNotifyStatus(uuid string, notifyStatus int64) error {
	return g.DB.Model(model.SpikeTx{}).Where("uuid = ?", uuid).Update("notify_status", notifyStatus).Error
}

func (g *GormAccessor) UpdateRechargeTxNotifyStatus(txHash string, notifyStatus int64) error {
	return g.DB.Model(model.SpikeTx{}).Where("tx_hash = ?", txHash).Update("notify_status", notifyStatus).Error
}
