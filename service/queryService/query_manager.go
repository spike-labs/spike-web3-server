package queryService

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/model"
	"spike-frame/response"
	"spike-frame/util"
	"time"
)

const nftListDuration = 20 * time.Minute

type QueryManager struct {
	sched       *Scheduler
	network     string
	redisClient *redis.Client
}

func NewQueryManager() *QueryManager {
	client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		panic("eth client dial err")
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

func (qm *QueryManager) QueryNftType(ctx context.Context, walletAddr string) ([]response.NftType, error) {
	value, isNil, err := util.GetStringFromRedis(walletAddr+constant.NFTTYPESUFFIX, qm.redisClient)
	if err != nil {
		return []response.NftType{}, err
	}
	if isNil {
		err = qm.queryNft(ctx, walletAddr)
		if err != nil {
			return []response.NftType{}, err
		}
	}

	value, isNil, err = util.GetStringFromRedis(walletAddr+constant.NFTTYPESUFFIX, qm.redisClient)
	if err != nil {
		return []response.NftType{}, err
	}
	var nftRes []response.NftType
	err = json.Unmarshal([]byte(value), &nftRes)
	if err != nil {
		return []response.NftType{}, err
	}
	return nftRes, nil
}

func (qm *QueryManager) QueryNftList(ctx context.Context, walletAddr string, assetType string) ([]response.NftResult, error) {

	value, isNil, err := util.GetStringFromRedis(walletAddr+constant.NFTLISTSUFFIX+assetType, qm.redisClient)
	if err != nil {
		return []response.NftResult{}, err
	}
	if isNil {
		err = qm.queryNft(ctx, walletAddr)
		if err != nil {
			return []response.NftResult{}, err
		}
	}

	value, isNil, err = util.GetStringFromRedis(walletAddr+constant.NFTLISTSUFFIX+assetType, qm.redisClient)
	if err != nil {
		return []response.NftResult{}, err
	}
	var nftRes []response.NftResult
	err = json.Unmarshal([]byte(value), &nftRes)
	if err != nil {
		return []response.NftResult{}, err
	}
	return nftRes, nil
}

func (qm *QueryManager) queryNft(ctx context.Context, walletAddr string) error {
	nftResults := make([]response.NftResult, 0)
	err := qm.sched.Schedule(ctx, model.NftQuery, func(ctx context.Context) error {
		nftResults, err := QueryWalletNft("", walletAddr, qm.network, nftResults)
		log.Infof("query wallet result : %+v", nftResults)
		return err
	})
	if err != nil {
		return err
	}
	_, err = qm.handleNftData(walletAddr, nftResults)
	return err
}

func (qm *QueryManager) QueryNativeRecord(ctx context.Context, walletAddr string) (response.BscResult, error) {
	value, isNil, err := util.GetStringFromRedis(walletAddr+constant.NATIVETXRECORDSUFFIX, qm.redisClient)
	if err != nil {
		return response.BscResult{}, err
	}

	if isNil {
		err = qm.queryNativeTxRecord(ctx, walletAddr)
		if err != nil {
			return response.BscResult{}, err
		}
	}

	value, isNil, err = util.GetStringFromRedis(walletAddr+constant.NATIVETXRECORDSUFFIX, qm.redisClient)
	if err != nil {
		return response.BscResult{}, err
	}
	var nativeRes response.BscResult
	err = json.Unmarshal([]byte(value), &nativeRes)
	if err != nil {
		return response.BscResult{}, err
	}
	return nativeRes, nil
}

func (qm *QueryManager) QueryERC20TxRecord(ctx context.Context, walletAddr string, contractAddr string) (response.BscResult, error) {

	value, isNil, err := util.GetStringFromRedis(walletAddr+contractAddr+constant.ERC20TXRECORDSUFFIX, qm.redisClient)
	if err != nil {
		return response.BscResult{}, err
	}
	if isNil {
		err = qm.queryERC20TxRecord(ctx, walletAddr, contractAddr)
		if err != nil {
			return response.BscResult{}, err
		}
	}

	value, isNil, err = util.GetStringFromRedis(walletAddr+contractAddr+constant.ERC20TXRECORDSUFFIX, qm.redisClient)
	if err != nil {
		return response.BscResult{}, err
	}
	var erc20Res response.BscResult
	err = json.Unmarshal([]byte(value), &erc20Res)
	if err != nil {
		return response.BscResult{}, err
	}
	return erc20Res, nil
}

func (qm *QueryManager) queryERC20TxRecord(ctx context.Context, walletAddr string, contractAddr string) error {
	bscResults := response.BscResult{}
	err := qm.sched.Schedule(ctx, model.Erc20TxRecordQuery, func(ctx context.Context) error {
		bscResults, err := queryERC20TxRecord(contractAddr, walletAddr)
		log.Infof("query er20 tx record result : %+v", bscResults)
		return err
	})
	if err != nil {
		return err
	}
	_, err = qm.handleERC20TxRecordData(walletAddr, contractAddr, bscResults)
	return err
}

func (qm *QueryManager) queryNativeTxRecord(ctx context.Context, walletAddr string) error {
	bscResults := response.BscResult{}
	err := qm.sched.Schedule(ctx, model.NativeTxRecordQuery, func(ctx context.Context) error {
		bscResults, err := queryNativeTxRecord(walletAddr)
		log.Infof("query native tx record result : %+v", bscResults)
		return err
	})
	if err != nil {
		return err
	}
	_, err = qm.handleNativeTxRecordData(walletAddr, bscResults)
	return err
}
