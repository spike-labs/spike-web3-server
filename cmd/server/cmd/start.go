package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spike-engine/spike-web3-server/chain"
	"github.com/spike-engine/spike-web3-server/initialize"
)

func StartCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start Spike Web3 Server",
		Long:  `According to the configuration item of config.toml, start this project`,
		Run: func(cmd *cobra.Command, args []string) {
			chain.NewBscListener()
			initialize.RunServer()
		},
	}

	return cmd
}
