package game

import "spike-frame/model"

type TxTracker interface {
	SaveTxCb(tx model.SpikeTx) error
	RecordTxHash(uuidList []string, txHash string, txStatus int) error
	QueryGameCb(txHash string) ([]model.SpikeTx, error)
	UpdateTxStatus(txHash string, txStatus int, payTime int64) error
	UpdateTxNotifyStatus(orderId string, notifyStatus int64) error
	QueryNotNotifyTx(notNotifyStatus int) ([]model.SpikeTx, error)
}
