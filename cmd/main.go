package main

import (
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/cmd/server/cmd"
)

func main() {
	logger.SetLogLevel("*", "INFO")
	cmd.Execute()
}
