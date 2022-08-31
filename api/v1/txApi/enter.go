package txApi

import "github.com/gin-gonic/gin"

type TxGroup struct {
}

func NewTxGroup() TxGroup {
	return TxGroup{}
}

func (txGroup *TxGroup) InitTxGroup(g *gin.RouterGroup) {
	g.Use()
	//todo
}
