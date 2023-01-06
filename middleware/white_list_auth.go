package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/util"
)

func WhiteListAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := util.GetRealIp(c.Request)
		log.Infof("clientIp : %s", clientIp)
		if util.IsLocalIp(clientIp) {
			c.Next()
			return
		}

		for _, ip := range config.Cfg.TxApiWhiteList.IpList {
			if ip == clientIp {
				c.Next()
				return
			}
		}
		response.FailWithMessage("Yor are not in whiteList", c)
		c.Abort()
	}
}
