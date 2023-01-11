package sign

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"

	"github.com/spike-engine/spike-web3-server/cache"
	chain "github.com/spike-engine/spike-web3-server/chain/abi"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/model"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/util"
)

type WorkerID uuid.UUID // worker session UUID

func (w WorkerID) String() string {
	return uuid.UUID(w).String()
}

type Worker interface {
	WorkerCalls
	GetInfo() *WorkerInfo
	Lock() bool
	UnLock()
	SignatureTransaction(*types.Transaction) (*types.Transaction, error)
	AddTaskNum()
}

type WorkerCalls interface {
	BatchMint(reqs []model.BatchMintReq) ([]string, string, error)
	WithdrawToken(reqs []model.WithdrawTokenReq) ([]string, string, error)
	WithdrawNFT(reqs []model.WithdrawNFTReq) ([]string, string, error)
}

type AllRoundWorker struct {
	BscClient  *ethclient.Client
	httpClient *resty.Client
	rdb        *redis.Client

	info     *WorkerInfo
	rLK      sync.Mutex
	nftABI   abi.ABI
	vaultABI abi.ABI
}

type WorkerInfo struct {
	walletAddress string
	serverUrl     string
	TaskNum       uint32
}

type UnSignTX struct {
	Tx *types.Transaction
}

func NewAllRoundWorker(workers config.SignWorker) (*AllRoundWorker, error) {

	bscClient, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		return nil, err
	}

	nftAbi, err := abi.JSON(strings.NewReader(chain.ERC721ContractABI))
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
		walletAddress: workers.WalletAddress,
		serverUrl:     workers.ServerUrl,
	}

	worker := &AllRoundWorker{
		BscClient:  bscClient,
		httpClient: resty.New(),
		rdb:        cache.RedisClient,
		info:       info,
		nftABI:     nftAbi,
		vaultABI:   vaultAbi,
	}

	CNonce, err := worker.GetCNonce()
	if err != nil {
		return nil, err
	}

	worker.rdb.Set(context.Background(), info.walletAddress+constant.NONCE, int64(CNonce), 0)

	return worker, nil
}

func (w *AllRoundWorker) GetInfo() *WorkerInfo {
	return w.info
}

func (w *AllRoundWorker) Lock() bool {
	lock, err := util.Lock(w.GetInfo().walletAddress, 1, 30*time.Second, w.rdb)
	if err != nil {
		log.Error("===Spike log:", err)
		return false
	}
	return lock
}

func (w *AllRoundWorker) UnLock() {
	err := util.UnLock(w.GetInfo().walletAddress, w.rdb)
	if err != nil {
		log.Error("===Spike log:", err)
		return
	}
	return
}

func (w *AllRoundWorker) IncrNonce() {
	w.rdb.Incr(context.Background(), w.GetInfo().walletAddress+constant.NONCE)
}

func (w *AllRoundWorker) GetRNonce() (uint64, error) {
	nonce, err := w.rdb.Get(context.Background(), w.GetInfo().walletAddress+constant.NONCE).Int64()
	if err != nil {
		return 0, err
	}
	return uint64(nonce), nil
}

func (w *AllRoundWorker) GetCNonce() (uint64, error) {
	nonce, err := w.BscClient.PendingNonceAt(context.Background(), common.HexToAddress(w.GetInfo().walletAddress))
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func (w *AllRoundWorker) AddTaskNum() {
	w.info.TaskNum++
}

func (w *AllRoundWorker) SignatureTransaction(unSignTX *types.Transaction) (*types.Transaction, error) {
	var res response.Response
	var signedTransaction types.Transaction

	resp, err := w.httpClient.R().
		SetHeader("Accept", "application/json").
		SetBody(&UnSignTX{Tx: unSignTX}).
		Post(w.GetInfo().serverUrl)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(res.Data.(string)), &signedTransaction)
	if err != nil {
		return nil, err
	}
	return &signedTransaction, nil
}

func (w *AllRoundWorker) BatchMint(reqs []model.BatchMintReq) ([]string, string, error) {
	w.rLK.Lock()
	defer w.rLK.Unlock()

	rNonce, err := w.GetRNonce()
	if err != nil {
		return nil, "", err
	}

	CNonce, err := w.GetCNonce()
	if err != nil {
		return nil, "", err
	}

	if rNonce != CNonce {
		return nil, "", errors.New("nonce error")
	}

	uuids := make([]string, 0)
	tokenIds := make([]*big.Int, 0)
	tokenUris := make([]string, 0)
	for i := range reqs {
		uuids = append(uuids, reqs[i].Uuid)
		tokenIds = append(tokenIds, big.NewInt(reqs[i].TokenID))
		tokenUris = append(tokenUris, reqs[i].TokenURI)
	}

	log.Infof("===Spike log : uuids:%v ;tokenids:%v ; tokenUris: %v", uuids, tokenIds, tokenUris)

	inputData, err := w.nftABI.Pack("batchMint0", tokenIds, common.HexToAddress(config.Cfg.Contract.GameVaultAddress), tokenUris)
	if err != nil {
		return nil, "", err
	}

	unSignTransaction, err := util.NewSpikeTx(common.HexToAddress(w.info.walletAddress), reqs[0].NFTAddress, inputData, CNonce, w.BscClient).ConstructionTransaction()
	if err != nil {
		return nil, "", err
	}

	signedTransaction, err := w.SignatureTransaction(unSignTransaction)
	if err != nil {
		return nil, "", err
	}

	err = w.BscClient.SendTransaction(context.Background(), signedTransaction)
	if err != nil {
		return nil, "", err
	}

	w.IncrNonce()

	return uuids, signedTransaction.Hash().String(), nil
}

func (w *AllRoundWorker) WithdrawToken(reqs []model.WithdrawTokenReq) ([]string, string, error) {
	w.rLK.Lock()
	defer w.rLK.Unlock()

	rNonce, err := w.GetRNonce()
	if err != nil {
		return nil, "", err
	}

	CNonce, err := w.GetCNonce()
	if err != nil {
		return nil, "", err
	}

	if rNonce != CNonce {
		return nil, "", errors.New("nonce error")
	}

	uuids := make([]string, 0)
	toAddrs := make([]common.Address, 0)
	tokenAddrs := make([]common.Address, 0)
	amounts := make([]*big.Int, 0)
	for i := range reqs {
		uuids = append(uuids, reqs[i].Uuid)
		amounts = append(amounts, util.ToWei(reqs[i].Amount, 18))
		toAddrs = append(toAddrs, reqs[i].ToAddress)
		tokenAddrs = append(tokenAddrs, reqs[i].TokenAddress)
	}

	log.Infof("===Spike log : uuids:%v ;toAddrs:%v ; tokenAddrs: %v ; amounts: %v", uuids, toAddrs, tokenAddrs, amounts)

	inputData, err := w.vaultABI.Pack("batchWithdraw0", tokenAddrs, toAddrs, amounts)
	if err != nil {
		return nil, "", err
	}

	unSignTransaction, err := util.NewSpikeTx(common.HexToAddress(w.info.walletAddress), config.Cfg.Contract.GameVaultAddress, inputData, CNonce, w.BscClient).ConstructionTransaction()
	if err != nil {
		return nil, "", err
	}

	signedTransaction, err := w.SignatureTransaction(unSignTransaction)
	if err != nil {
		return nil, "", err
	}

	err = w.BscClient.SendTransaction(context.Background(), signedTransaction)
	if err != nil {
		return nil, "", err
	}

	w.IncrNonce()

	return uuids, signedTransaction.Hash().String(), nil
}

func (w *AllRoundWorker) WithdrawNFT(reqs []model.WithdrawNFTReq) ([]string, string, error) {
	w.rLK.Lock()
	defer w.rLK.Unlock()

	rNonce, err := w.GetRNonce()
	if err != nil {
		return nil, "", err
	}

	CNonce, err := w.GetCNonce()
	if err != nil {
		return nil, "", err
	}

	if rNonce != CNonce {
		return nil, "", errors.New("nonce error")
	}

	uuids := make([]string, 0)
	toAddrs := make([]common.Address, 0)
	tokenAddrs := make([]common.Address, 0)
	tokenIds := make([]*big.Int, 0)
	for i := range reqs {
		uuids = append(uuids, reqs[i].Uuid)
		tokenIds = append(tokenIds, big.NewInt(reqs[i].TokenId))
		toAddrs = append(toAddrs, reqs[i].ToAddress)
		tokenAddrs = append(tokenAddrs, reqs[i].TokenAddress)
	}

	log.Infof("===Spike log : uuids:%v ;toAddrs:%v ; tokenAddrs: %v ; tokenIds: %v", uuids, toAddrs, tokenAddrs, tokenIds)

	inputData, err := w.vaultABI.Pack("batchWithdrawNFT", tokenAddrs, toAddrs, tokenIds)
	if err != nil {
		return nil, "", err
	}

	unSignTransaction, err := util.NewSpikeTx(common.HexToAddress(w.info.walletAddress), config.Cfg.Contract.GameVaultAddress, inputData, CNonce, w.BscClient).ConstructionTransaction()
	if err != nil {
		return nil, "", err
	}

	signedTransaction, err := w.SignatureTransaction(unSignTransaction)
	if err != nil {
		return nil, "", err
	}

	err = w.BscClient.SendTransaction(context.Background(), signedTransaction)
	if err != nil {
		return nil, "", err
	}

	w.IncrNonce()
	return uuids, signedTransaction.Hash().String(), nil
}
