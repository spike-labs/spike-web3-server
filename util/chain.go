package util

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spike-engine/spike-web3-server/config"
)

func QueryBlockHeight() (uint64, error) {
	client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		return 0, err
	}
	blockNum, err := client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}
	return blockNum, nil
}
