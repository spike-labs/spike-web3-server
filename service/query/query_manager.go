package query

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/spike-engine/spike-web3-server/cache"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/model"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/util"
	"time"
)

const nftListDuration = 20 * time.Minute

var QurManager *QueryManager

type QueryManager struct {
	sched       *Scheduler
	network     string
	redisClient *redis.Client
}

func NewQueryManager() *QueryManager {
	client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		panic("eth client dial err")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		panic("query chainId err")
	}
	var network string
	log.Infof("chainId: %s", chainId.String())
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
		redisClient: cache.RedisClient,
	}
}

func (qm *QueryManager) NftList(gameContractAddr, walletAddr string) ([]model.CacheData, error) {
	nftResults := make([]response.NftResult, 0)
	results, err := QueryWalletNft("", gameContractAddr, walletAddr, qm.network, nftResults)
	if err != nil {
		return []model.CacheData{}, err
	}
	data := util.ConvertNftResult(gameContractAddr, results)
	dataList := util.ParseMetadata(data)
	return dataList, nil
}

func (qm *QueryManager) QueryNftType(ctx context.Context, gameContractAddr, walletAddr string) ([]response.NftType, error) {
	value, isNil, err := util.GetStringFromRedis(walletAddr+gameContractAddr+constant.NFTTYPESUFFIX, qm.redisClient)
	if err != nil {
		return []response.NftType{}, err
	}
	if isNil {
		err = qm.queryNft(ctx, gameContractAddr, walletAddr)
		if err != nil {
			return []response.NftType{}, err
		}
	}

	value, isNil, err = util.GetStringFromRedis(walletAddr+gameContractAddr+constant.NFTTYPESUFFIX, qm.redisClient)
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

func (qm *QueryManager) QueryNftList(ctx context.Context, gameContractAddr, walletAddr, assetType string) ([]model.CacheData, error) {

	value, isNil, err := util.GetStringFromRedis(walletAddr+gameContractAddr+constant.NFTLISTSUFFIX+assetType, qm.redisClient)
	if err != nil {
		return []model.CacheData{}, err
	}
	if isNil {
		err = qm.queryNft(ctx, gameContractAddr, walletAddr)
		if err != nil {
			return []model.CacheData{}, err
		}
	}

	value, isNil, err = util.GetStringFromRedis(walletAddr+gameContractAddr+constant.NFTLISTSUFFIX+assetType, qm.redisClient)
	if err != nil {
		return []model.CacheData{}, err
	}
	var nftRes []model.CacheData
	err = json.Unmarshal([]byte(value), &nftRes)
	if err != nil {
		return []model.CacheData{}, err
	}
	return nftRes, nil
}

func (qm *QueryManager) queryNft(ctx context.Context, gameNftAddr, walletAddr string) error {
	nftResults := make([]response.NftResult, 0)
	err := qm.sched.Schedule(ctx, model.NftQuery, func(ctx context.Context) error {
		results, err := QueryWalletNft("", gameNftAddr, walletAddr, qm.network, nftResults)
		log.Infof("query wallet result : %+v", results)
		nftResults = results
		return err
	})
	if err != nil {
		return err
	}
	_, err = qm.handleNftData(gameNftAddr, walletAddr, nftResults)
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
		results, err := queryERC20TxRecord(contractAddr, walletAddr)
		log.Infof("query er20 tx record result : %+v", results)
		bscResults = results
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
		results, err := queryNativeTxRecord(walletAddr)
		log.Infof("query native tx record result : %+v", results)
		bscResults = results
		return err
	})
	if err != nil {
		return err
	}
	_, err = qm.handleNativeTxRecordData(walletAddr, bscResults)
	return err
}
