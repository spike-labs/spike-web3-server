package chain

type TxFilter interface {
	Accept(fromAddr, toAddr string) bool
}
