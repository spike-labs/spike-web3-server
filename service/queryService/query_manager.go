package queryService

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/model"
	"spike-frame/response"
)

type QueryManager struct {
	sched       *Scheduler
	network     string
	redisClient *redis.Client
}

func NewQueryManager() *QueryManager {
	client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		log.Error("eth client dial err : ", err)
		return nil
	}
	chainId, err := client.ChainID(context.Background())
	var network string
	switch chainId.String() {
	case "56":
		network = "bsc"
	case "97":
		network = "bsc testnet"
	default:
		panic("not expected chainId")
	}

	return &QueryManager{
		network:     network,
		sched:       NewScheduler(),
		redisClient: constant.RedisClient,
	}
}

func (qm *QueryManager) QueryNftList(ctx context.Context, walletAddr string) ([]response.NftResult, error) {
	//todo cache
	nftResults := make([]response.NftResult, 0)
	err := qm.sched.Schedule(ctx, model.NftQuery, func(ctx context.Context) error {
		nftResults, _ = QueryWalletNft("", walletAddr, qm.network, nftResults)
		return nil
	})
	if err != nil {
		return nftResults, err
	}
	return nftResults, nil
}
