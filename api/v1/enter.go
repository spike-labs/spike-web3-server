package v1

import (
	logger "github.com/ipfs/go-log"
	"spike-frame/api/v1/queryApi"
	"spike-frame/api/v1/txApi"
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