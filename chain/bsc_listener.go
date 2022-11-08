package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	logger "github.com/ipfs/go-log"
	"github.com/spike-engine/spike-web3-server/cache"
	chain "github.com/spike-engine/spike-web3-server/chain/abi"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/dao"
	"github.com/spike-engine/spike-web3-server/game"
	"github.com/spike-engine/spike-web3-server/util"
	"math/big"
	"sync"
	"time"
)

var log = logger.Logger("chain")

const (
	BLOCKNUM = "blockNum"
)

type ErrMsg struct {
	contractAddr string
	from         *big.Int
	to           *big.Int
}

type BscListener struct {
	network     string
	ec          *ethclient.Client
	rc          *redis.Client
	l           map[string]Listener
	errorHandle chan ErrMsg
}

func NewBscListener() (*BscListener, error) {
	log.Infof("bsc listener start")
	bl := &BscListener{}
	client, err := ethclient.Dial(config.Cfg.Chain.WsNodeAddress)
	if err != nil {
		log.Error("eth client dial err : ", err)
		return nil, err
	}

	errorHandle := make(chan ErrMsg, 10)
	bl.errorHandle = errorHandle
	bl.ec = client

	cbManger := game.NewCbManager(dao.DbAccessor)
	go cbManger.Run()

	l := make(map[string]Listener)
	l[constant.EmptyAddress] = newBNBListener(bl.ec, errorHandle)
	gameVaultChan := make(util.DataChannel, 10)
	l[config.Cfg.Contract.GameVaultAddress] = newERC20Listener(config.Cfg.Contract.GameVaultAddress, bl.ec, gameVaultChan, util.GetABI(chain.GameVaultABI), errorHandle)
	util.Eb.Subscribe(constant.NewBlockTopic, gameVaultChan)

	for _, contractAddr := range config.Cfg.Contract.ERC20ContractAddress {
		erc20TokenChan := make(util.DataChannel, 10)
		l[contractAddr] = newERC20Listener(contractAddr, bl.ec, erc20TokenChan, util.GetABI(chain.ERC20ContractABI), errorHandle)
		util.Eb.Subscribe(constant.NewBlockTopic, erc20TokenChan)
	}

	for _, contractAddr := range config.Cfg.Contract.NftContractAddress {
		erc721TokenChan := make(util.DataChannel, 10)
		l[contractAddr] = newERC721Listener(contractAddr, bl.ec, erc721TokenChan, util.GetABI(chain.ERC721ContractABI), errorHandle)
		util.Eb.Subscribe(constant.NewBlockTopic, erc721TokenChan)
	}
	bl.l = l
	go bl.Run()
	return bl, nil
}

func (bl *BscListener) Run() {
	go bl.handleError()
	//sync
	for {
		nowBlockNum, err := bl.ec.BlockNumber(context.Background())
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			log.Error("query now bsc_blockNum err :", err)
			continue
		}
		cacheBlockNum, isNil, err := util.GetIntFromRedis(BLOCKNUM+config.Cfg.System.MachineId, cache.RedisClient)
		if err != nil {
			log.Errorf("query redis err : %v", err)
			return
		}
		if isNil {
			log.Infof("blockNum is not exist")
			util.SetFromRedis(BLOCKNUM+config.Cfg.System.MachineId, nowBlockNum-constant.BlockConfirmHeight, 0, cache.RedisClient)
			break
		}

		if cacheBlockNum >= int64(nowBlockNum)-constant.BlockConfirmHeight {
			log.Infof("sync done")
			break
		}
		var wg sync.WaitGroup
		for _, listener := range bl.l {
			wg.Add(1)
			go func(l Listener) {
				defer wg.Done()
				l.handlePastBlock(big.NewInt(cacheBlockNum+1), big.NewInt(int64(nowBlockNum-constant.BlockConfirmHeight)))
			}(listener)
		}
		wg.Wait()
		util.SetFromRedis(BLOCKNUM+config.Cfg.System.MachineId, nowBlockNum-constant.BlockConfirmHeight, 0, cache.RedisClient)
	}

	for _, listener := range bl.l {
		go func(l Listener) {
			l.run()
		}(listener)
	}
}

func (bl *BscListener) handleError() {
	for {
		select {
		case msg := <-bl.errorHandle:
			log.Infof("handle err ,type : %s, from : %d, to : %d", msg.contractAddr, msg.from.Int64(), msg.to.Int64())
			if _, ok := bl.l[msg.contractAddr]; ok {
				time.Sleep(200 * time.Millisecond)
				bl.l[msg.contractAddr].handlePastBlock(msg.from, msg.to)
			}
		}
	}
}
