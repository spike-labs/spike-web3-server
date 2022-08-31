package v1

import (
	"spike-frame/api/v1/queryApi"
	"spike-frame/api/v1/txApi"
)

type RouterGroup struct {
	QueryGroup queryApi.QueryGroup
	TxGroup    txApi.TxGroup
}

func NewRouterGroup() RouterGroup {
	return RouterGroup{
		QueryGroup: queryApi.NewQueryApiGroup(),
		TxGroup:    txApi.NewTxGroup(),
	}
}
