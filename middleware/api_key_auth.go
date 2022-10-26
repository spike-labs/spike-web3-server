package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spike-engine/spike-web3-server/global"
	"github.com/spike-engine/spike-web3-server/response"
)

func ApiKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params authParams
		if err := c.ShouldBindHeader(&params); err != nil {
			response.FailWithMessage("header: api_key is required", c)
			c.Abort()
		} else {
			res, _ := global.RedisClient.SIsMember(context.Background(), "api_key", params.APIKey).Result()
			if res {
				c.Next()
			} else {
				response.FailWithMessage("header: api_key doesn't exist", c)
				c.Abort()
			}
		}
	}
}

type authParams struct {
	APIKey string `header:"api_key" binding:"required"`
}
