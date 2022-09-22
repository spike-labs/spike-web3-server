package game

import "github.com/spike-engine/spike-web3-server/model"

type TxTracker interface {
	SaveTxCb(tx model.SpikeTx) error
	RecordTxHash(uuidList []string, txHash string, txStatus int) error
	QueryGameCb(txHash string, notifyStatus int) ([]model.SpikeTx, error)
	UpdateTxStatus(txHash string, txStatus int, payTime int64) error
	UpdateTxNotifyStatus(orderId string, notifyStatus int) error
	QueryNotNotifyTx(notNotifyStatus int) ([]model.SpikeTx, error)
}
