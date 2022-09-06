package chain

import "math/big"

type Listener interface {
	TxFilter
	run()
	handlePastBlock(fromBlock, toBlock *big.Int) error
}
