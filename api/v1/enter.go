package v1

import (
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/api/v1/query"
	"github.com/spike-engine/spike-web3-server/api/v1/tx"
)

var log = logger.Logger("api")

type RouterGroup struct {
	QueryGroup queryApi.QueryGroup
	TxGroup    txApi.TxGroup
}

func NewRouterGroup() (RouterGroup, error) {
	txGroup, err := txApi.NewTxGroup()
	if err != nil {
		log.Error("===Spike log:", err)
		return RouterGroup{}, err
	}
	return RouterGroup{
		QueryGroup: queryApi.NewQueryApiGroup(),
		TxGroup:    txGroup,
	}, nil
}
