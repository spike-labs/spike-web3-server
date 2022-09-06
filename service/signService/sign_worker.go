package signService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-resty/resty/v2"
	"math/big"
	chain "spike-frame/chain/abi"
	"spike-frame/config"
	"spike-frame/model"
	"spike-frame/response"
	"spike-frame/util"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type WorkerID uuid.UUID // worker session UUID

func (w WorkerID) String() string {
	return uuid.UUID(w).String()
}

type Worker interface {
	WorkerCalls
	GetInfo() *WorkerInfo
	Lock() bool
	UnLock() uint64
}

type WorkerCalls interface {
	BatchMint(reqs *model.BatchMintQueue) error
	WithdrawToken(reqs *model.WithdrawTokenQueue) error
	WithdrawNFT(reqs *model.WithdrawNFTQueue) error
}

type AllRoundWorker struct {
	BscClient  *ethclient.Client
	httpClient *resty.Client
	rdb        *redis.Client
	info       *WorkerInfo
	nLK        sync.Mutex
	nftABI     abi.ABI
	vaultABI   abi.ABI
}

type WorkerInfo struct {
	walletAddress common.Address
	serverUrl     string
	TaskNum       uint32
}

func NewAllRoundWorker() (*AllRoundWorker, error) {

	bscClient, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		return nil, err
	}

	nftAbi, err := abi.JSON(strings.NewReader(chain.GameNftABI))
	if err != nil {
		return nil, err
	}

	vaultAbi, err := abi.JSON(strings.NewReader(chain.GameVaultABI))
	if err != nil {
		return nil, err
	}

	if len(config.Cfg.SignWorkers) == 0 {
		return nil, errors.New("worker info can not be null")
	}

	info := &WorkerInfo{
		walletAddress: common.HexToAddress(config.Cfg.SignWorkers[0].WalletAddress),
		serverUrl:     config.Cfg.SignWorkers[0].ServerUrl,
	}

	config.Cfg.SignWorkers = config.Cfg.SignWorkers[1:]

	nonce, err := bscClient.PendingNonceAt(context.Background(), info.walletAddress)
	if err != nil {
		return nil, err
	}

	worker := &AllRoundWorker{
		BscClient:  bscClient,
		httpClient: resty.New(),
		rdb: redis.NewClient(&redis.Options{
			Addr:     config.Cfg.Redis.Address,
			Password: config.Cfg.Redis.Password,
			DB:       0,
		}),
		info:     info,
		nftABI:   nftAbi,
		vaultABI: vaultAbi,
	}

	worker.rdb.IncrBy(context.Background(), info.walletAddress.String()+":nonce", int64(nonce))

	return worker, nil
}

func (w *AllRoundWorker) GetInfo() *WorkerInfo {
	return w.info
}

func (w *AllRoundWorker) Lock() bool {
	w.nLK.Lock()
	defer w.nLK.Unlock()
	bool, err := w.rdb.SetNX(context.Background(), w.GetInfo().walletAddress.String(), 1, 10*time.Second).Result()
	if err != nil {
		log.Error("===Spike log:", err)
		return false
	}
	return bool
}

func (w *AllRoundWorker) UnLock() uint64 {
	nums, err := w.rdb.Del(context.Background(), w.GetInfo().walletAddress.String()).Result()
	if err != nil {
		log.Error("===Spike log:", err)
		return 0
	}
	return uint64(nums)
}

func (w *AllRoundWorker) IncrNonce() {
	w.rdb.Incr(context.Background(), w.GetInfo().walletAddress.String()+":nonce")
}

func (w *AllRoundWorker) GetRNonce() (uint64, error) {
	nonce, err := w.rdb.Get(context.Background(), w.GetInfo().walletAddress.String()+":nonce").Int64()
	if err != nil {
		return 0, err
	}
	return uint64(nonce), nil
}

func (w *AllRoundWorker) GetCNonce() (uint64, error) {
	nonce, err := w.BscClient.PendingNonceAt(context.Background(), w.GetInfo().walletAddress)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func (w *AllRoundWorker) BatchMint(queue *model.BatchMintQueue) error {
	w.Lock()
	defer w.UnLock()
	var res response.Response
	var signedTransaction types.Transaction

	rNonce, err := w.GetRNonce()
	if err != nil {
		return err
	}

	CNonce, err := w.GetCNonce()
	if err != nil {
		return err
	}

	if rNonce != CNonce {
		return errors.New("nonce error")
	}

	tokenIds := make([]*big.Int, 0)
	tokenUris := make([]string, 0)
	for i := range queue.Reqs {
		tokenIds = append(tokenIds, big.NewInt(queue.Reqs[i].TokenID))
		tokenUris = append(tokenUris, queue.Reqs[i].TokenURI)
	}

	inputData, err := w.nftABI.Pack("batchMint0", tokenIds, config.Cfg.Contract.GameVaultAddress, tokenUris)
	if err != nil {
		return err
	}

	spikeTx := &util.SpikeTx{
		Data:      inputData,
		To:        config.Cfg.Contract.GameNftAddress,
		BscClient: w.BscClient,
		From:      w.info.walletAddress,
		Nonce:     CNonce,
	}

	transaction, err := spikeTx.ConstructionTransaction()
	if err != nil {
		return err
	}

	resp, err := w.httpClient.R().
		SetHeader("Accept", "application/json").
		SetBody(transaction).
		Post(w.GetInfo().serverUrl)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(res.Data.(string)), &signedTransaction)
	if err != nil {
		fmt.Println(err)
	}

	err = w.BscClient.SendTransaction(context.Background(), &signedTransaction)
	if err != nil {
		return err
	}

	w.IncrNonce()

	return nil
}

func (w *AllRoundWorker) WithdrawToken(queue *model.WithdrawTokenQueue) error {
	w.Lock()
	defer w.UnLock()
	var res response.Response
	var signedTransaction types.Transaction

	rNonce, err := w.GetRNonce()
	if err != nil {
		return err
	}

	CNonce, err := w.GetCNonce()
	if err != nil {
		return err
	}

	if rNonce != CNonce {
		return errors.New("nonce error")
	}

	toAddrs := make([]common.Address, 0)
	tokenAddrs := make([]common.Address, 0)
	amounts := make([]*big.Int, 0)
	for i := range queue.Reqs {
		amounts = append(amounts, big.NewInt(queue.Reqs[i].Amount))
		toAddrs = append(toAddrs, queue.Reqs[i].ToAddress)
		tokenAddrs = append(tokenAddrs, queue.Reqs[i].TokenAddress)
	}

	inputData, err := w.vaultABI.Pack("batchWithdraw0", tokenAddrs, toAddrs, amounts)
	if err != nil {
		return err
	}

	spikeTx := &util.SpikeTx{
		Data:      inputData,
		To:        config.Cfg.Contract.GameVaultAddress,
		BscClient: w.BscClient,
		From:      w.info.walletAddress,
		Nonce:     CNonce,
	}

	transaction, err := spikeTx.ConstructionTransaction()
	if err != nil {
		return err
	}

	resp, err := w.httpClient.R().
		SetHeader("Accept", "application/json").
		SetBody(transaction).
		Post(w.GetInfo().serverUrl)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(res.Data.(string)), &signedTransaction)
	if err != nil {
		fmt.Println(err)
	}

	err = w.BscClient.SendTransaction(context.Background(), transaction)
	if err != nil {
		return err
	}

	w.IncrNonce()

	return nil
}

func (w *AllRoundWorker) WithdrawNFT(queue *model.WithdrawNFTQueue) error {
	w.Lock()
	defer w.UnLock()
	var res response.Response
	var signedTransaction types.Transaction

	rNonce, err := w.GetRNonce()
	if err != nil {
		return err
	}

	CNonce, err := w.GetCNonce()
	if err != nil {
		return err
	}

	if rNonce != CNonce {
		return errors.New("nonce error")
	}

	toAddrs := make([]common.Address, 0)
	tokenAddrs := make([]common.Address, 0)
	tokenIds := make([]*big.Int, 0)
	for i := range queue.Reqs {
		tokenIds = append(tokenIds, big.NewInt(queue.Reqs[i].TokenId))
		toAddrs = append(toAddrs, queue.Reqs[i].ToAddress)
		tokenAddrs = append(tokenAddrs, queue.Reqs[i].TokenAddress)
	}

	inputData, err := w.vaultABI.Pack("batchWithdrawNFT", tokenAddrs, toAddrs, tokenIds)
	if err != nil {
		return err
	}

	spikeTx := &util.SpikeTx{
		Data:      inputData,
		To:        config.Cfg.Contract.GameVaultAddress,
		BscClient: w.BscClient,
		From:      w.info.walletAddress,
		Nonce:     CNonce,
	}

	transaction, err := spikeTx.ConstructionTransaction()
	if err != nil {
		return err
	}

	resp, err := w.httpClient.R().
		SetHeader("Accept", "application/json").
		SetBody(transaction).
		Post(w.GetInfo().serverUrl)
	if err != nil {
		return err
	}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(res.Data.(string)), &signedTransaction)
	if err != nil {
		fmt.Println(err)
	}

	err = w.BscClient.SendTransaction(context.Background(), transaction)
	if err != nil {
		return err
	}

	w.IncrNonce()
	return nil
}
