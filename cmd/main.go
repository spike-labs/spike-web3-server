package main

import (
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/cmd/server/cmd"
	_ "github.com/spike-engine/spike-web3-server/docs"
)

// @title   Swagger Example API
// @version 0.0.1
// @description
// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       api_key
// @BasePath                   /
func main() {
	logger.SetLogLevel("*", "INFO")
	cmd.Execute()
}
