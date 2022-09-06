package chain

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	chain "spike-frame/chain/abi"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/model"
	"spike-frame/util"
	"strings"
)

func (e *ERC721Listener) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(constant.EmptyAddress) == strings.ToLower(fromAddr) {
		return true, constant.GAMENFT_TRANSFER
	}

	if strings.ToLower(config.Cfg.Contract.GameVaultAddress) == strings.ToLower(toAddr) {
		return true, constant.GAMENFT_IMPORT
	}
	return true, constant.GAMENFT_TRANSFER
}

type ERC721Listener struct {
	contractAddr   string
	tokenType      model.TokenType
	newBlockNotify util.DataChannel
	ec             *ethclient.Client
	abi            abi.ABI
	errorHandler   chan ErrMsg
}

func newERC721Listener(contractAddr string, tokenType model.TokenType, ec *ethclient.Client, newBlockNotify util.DataChannel, abi abi.ABI, errorHandler chan ErrMsg) *ERC721Listener {
	return &ERC721Listener{
		contractAddr,
		tokenType,
		newBlockNotify,
		ec,
		abi,
		errorHandler,
	}
}

func (al *ERC721Listener) run() {
	go al.NewEventFilter()
}

func (al *ERC721Listener) NewEventFilter() error {
	for {
		select {
		case de := <-al.newBlockNotify:
			height := de.Data.(*big.Int)
			al.handlePastBlock(height, height)
		}
	}
}

func (al *ERC721Listener) handlePastBlock(fromBlockNum, toBlockNum *big.Int) error {
	log.Infof("nft past event filter, fromBlock : %d, toBlock : %d ", fromBlockNum, toBlockNum)
	ethClient := al.ec
	contractAddress := common.HexToAddress(al.contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: fromBlockNum,
		ToBlock:   toBlockNum,
	}

	sub, err := ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		al.errorHandler <- ErrMsg{
			tp:   al.tokenType,
			from: fromBlockNum,
			to:   toBlockNum,
		}
		log.Errorf("nft subscribe event log, from: %d,to: %d,err : %+v", fromBlockNum.Int64(), toBlockNum.Int64(), err)
		return err
	}
	for _, l := range sub {
		switch l.Topics[0].String() {
		case util.EventSignHash(chain.TRANSFERTOPIC):
			msg := ErrMsg{
				tp:   al.tokenType,
				from: big.NewInt(int64(l.BlockNumber)),
				to:   big.NewInt(int64(l.BlockNumber)),
			}
			recp, err := al.ec.TransactionReceipt(context.Background(), l.TxHash)
			if err != nil {
				al.errorHandler <- msg
				log.Error("nft TransactionReceipt err : ", err)
				break
			}
			block, err := al.ec.BlockByNumber(context.Background(), big.NewInt(int64(l.BlockNumber)))
			if err != nil {
				al.errorHandler <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", l.BlockNumber, err)
				break
			}

			fromAddr := common.HexToAddress(l.Topics[1].Hex()).String()
			toAddr := common.HexToAddress(l.Topics[2].Hex()).String()
			_, txType := al.Accept(fromAddr, toAddr)
			//al.rc.Del(fromAddr + nftTypeSuffix)
			//al.rc.Del(toAddr + nftTypeSuffix)
			//al.rc.Del(fromAddr + Soul)
			//al.rc.Del(fromAddr + Soul_Tank)
			//al.rc.Del(toAddr + Soul_Tank)
			//al.rc.Del(toAddr + Soul)

			_ = model.SpikeTx{
				From:    fromAddr,
				To:      toAddr,
				TxType:  txType,
				TxHash:  l.TxHash.Hex(),
				Status:  recp.Status,
				PayTime: int64(block.Time() * 1000),
				TokenId: l.Topics[3].Big().Uint64(),
			}
		}
	}
	return nil
}
