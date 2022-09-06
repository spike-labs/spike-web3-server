package dao

import "spike-frame/model"

type TxTracker interface {
	SaveTxCb(orderId, uuid, txHash, from, to, amount, tokenId, cb string, txType, createTime int64) error
	RecordWithdrawTxHash(uuidList []string, txHash string, txStatus int64) error
	QueryGameCb(txHash string) ([]model.SpikeTx, error)
	UpdateTxStatus(txHash string, txStatus, payTime int64) error
	UpdateWithdrawTxNotifyStatus(uuid string, notifyStatus int64) error
	UpdateRechargeTxNotifyStatus(txHash string, notifyStatus int64) error
}
