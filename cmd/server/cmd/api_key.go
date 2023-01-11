package cmd

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spike-engine/spike-web3-server/cache"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/model"
)

func ApiKeyCommands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api_key",
		Short: "Manage the api_key",
	}

	cmd.AddCommand(
		AddApiKeyCommand())
	return cmd
}

func AddApiKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add a api_key",
		Args:  cobra.ExactArgs(1),
		RunE:  runAddCmd,
	}
	return cmd
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	apiKey := args[0]
	cache.RedisClient.SAdd(context.Background(), "api_key", apiKey)
	err := dao.DbAccessor.AddApiKey(model.ApiKey{
		Id:     uuid.New().String(),
		ApiKey: apiKey,
	})
	if err != nil {
		fmt.Println("add apiKey err : ", err)
		return err
	}
	fmt.Println("add apiKey success : ", apiKey)
	return nil
}
