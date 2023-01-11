package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spike-engine/spike-web3-server/cache"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/initialize"
	"github.com/spike-engine/spike-web3-server/service/query"
	"github.com/spike-engine/spike-web3-server/service/sign"
)

var (
	rootCmd *cobra.Command
)

func RootCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "web3 server",
		Short: "Spike Web3 server",
	}
	return cmd
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	config.Viper = config.InitViper()
	dao.GormClient = initialize.GormMysql()
	dao.DbAccessor = dao.NewGormAccessor(dao.GormClient)
	cache.RedisClient = cache.ConnectRedis()

	query.QurManager = query.NewQueryManager()
	sign.HwManager = sign.NewHWManager()

	rootCmd = RootCommands()
	rootCmd.AddCommand(
		StartCommand(),
		ApiKeyCommands())
}
