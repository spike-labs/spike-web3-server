package game

import "spike-frame/model"

type TxTracker interface {
	SaveTxCb(tx model.SpikeTx) error
	RecordTxHash(uuidList []string, txHash string, txStatus int) error
	QueryGameCb(txHash string) ([]model.SpikeTx, error)
	UpdateTxStatus(txHash string, txStatus int, payTime int64) error
	UpdateWithdrawTxNotifyStatus(uuid string, notifyStatus int64) error
	UpdateRechargeTxNotifyStatus(txHash string, notifyStatus int64) error
}
