package chain

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	logger "github.com/ipfs/go-log"
	"math/big"
	chain "spike-frame/chain/abi"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/model"
	"spike-frame/util"
	"sync"
	"time"
)

var log = logger.Logger("chain")

const (
	BLOCKNUM = "blockNum"
)

type ErrMsg struct {
	tp   model.TokenType
	from *big.Int
	to   *big.Int
}

type BscListener struct {
	network     string
	ec          *ethclient.Client
	rc          *redis.Client
	l           map[model.TokenType]Listener
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

	gameVaultChan := make(util.DataChannel, 10)
	governanceTokenChan := make(util.DataChannel, 10)
	usdcTokenChan := make(util.DataChannel, 10)
	gameTokenChan := make(util.DataChannel, 10)
	gameNftChan := make(util.DataChannel, 10)
	util.Eb.Subscribe(constant.NewBlockTopic, gameVaultChan)
	util.Eb.Subscribe(constant.NewBlockTopic, governanceTokenChan)
	util.Eb.Subscribe(constant.NewBlockTopic, usdcTokenChan)
	util.Eb.Subscribe(constant.NewBlockTopic, gameTokenChan)
	util.Eb.Subscribe(constant.NewBlockTopic, gameNftChan)

	l := make(map[model.TokenType]Listener)
	l[model.Bnb] = newBNBListener(bl.ec, errorHandle)
	l[model.Usdc] = newERC20Listener(config.Cfg.Contract.UsdcAddress, model.Usdc, bl.ec, usdcTokenChan, util.GetABI(chain.USDCContractABI), errorHandle)
	l[model.GovernanceToken] = newERC20Listener(config.Cfg.Contract.GovernanceTokenAddress, model.GovernanceToken, bl.ec, governanceTokenChan, util.GetABI(chain.GovernanceTokenABI), errorHandle)
	l[model.GameToken] = newERC20Listener(config.Cfg.Contract.GameTokenAddress, model.GameToken, bl.ec, gameTokenChan, util.GetABI(chain.GameTokenABI), errorHandle)
	l[model.GameVault] = newERC20Listener(config.Cfg.Contract.GameVaultAddress, model.GameVault, bl.ec, gameVaultChan, util.GetABI(chain.GameVaultABI), errorHandle)
	l[model.GameNft] = newERC721Listener(config.Cfg.Contract.GameNftAddress, model.GameNft, bl.ec, gameNftChan, util.GetABI(chain.GameNftABI), errorHandle)
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
			log.Error("query now bnb_blockNum err :", err)
			continue
		}
		cacheBlockNum, isNil, err := util.GetIntFromRedis(BLOCKNUM+config.Cfg.System.MachineId, constant.RedisClient)
		if err != nil {
			log.Errorf("query redis err : %v", err)
			return
		}
		if isNil {
			log.Infof("blockNum is not exist")
			util.SetFromRedis(BLOCKNUM+config.Cfg.System.MachineId, nowBlockNum-constant.BlockConfirmHeight, 0, constant.RedisClient)
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
		util.SetFromRedis(BLOCKNUM+config.Cfg.System.MachineId, nowBlockNum-constant.BlockConfirmHeight, 0, constant.RedisClient)
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
			log.Infof("handle err ,type : %s, from : %d, to : %d", msg.tp.String(), msg.from.Int64(), msg.to.Int64())
			if _, ok := bl.l[msg.tp]; ok {
				time.Sleep(200 * time.Millisecond)
				bl.l[msg.tp].handlePastBlock(msg.from, msg.to)
			}
		}
	}
}
