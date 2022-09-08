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

func (e *ERC20Listener) Accept(fromAddr, toAddr string) (bool, uint64) {
	if strings.ToLower(config.Cfg.Contract.GameVaultAddress) == strings.ToLower(toAddr) {
		switch e.tokenType {
		case model.Usdc:
			return true, constant.USDC_RECHARGE
		case model.GameToken:
			return true, constant.GAMETOKEN_RECHARGE
		case model.GovernanceToken:
			return true, constant.GOVERNANCETOKEN_RECHARGE
		}
	}

	if strings.ToLower(config.Cfg.Contract.GameVaultAddress) == strings.ToLower(fromAddr) {
		switch e.tokenType {
		case model.Usdc:
			return true, constant.USDC_WITHDRAW
		case model.GameToken:
			return true, constant.GAMETOKEN_WITHDRAW
		case model.GovernanceToken:
			return true, constant.GOVERNANCE_WITHDRAW
		case model.GameVault:
			return true, constant.NATIVE_WITHDRAW
		}
	}

	return false, constant.NOT_EXIST
}

type ERC20Listener struct {
	contractAddr   string
	tokenType      model.TokenType
	newBlockNotify util.DataChannel
	ec             *ethclient.Client
	abi            abi.ABI
	errorHandler   chan ErrMsg
}

func newERC20Listener(contractAddr string, tokenType model.TokenType, ec *ethclient.Client, newBlockNotify util.DataChannel, abi abi.ABI, errorHandler chan ErrMsg) *ERC20Listener {
	el := &ERC20Listener{
		contractAddr,
		tokenType,
		newBlockNotify,
		ec,
		abi,
		errorHandler,
	}
	return el
}

func (el *ERC20Listener) run() {
	go el.NewEventFilter(el.contractAddr)
}

func (el *ERC20Listener) NewEventFilter(contractAddr string) error {
	for {
		select {
		case de := <-el.newBlockNotify:
			height := de.Data.(*big.Int)
			el.handlePastBlock(height, height)
		}
	}
}

func (el *ERC20Listener) handlePastBlock(fromBlockNum, toBlockNum *big.Int) error {
	log.Infof("erc20 past event filter, type : %v, fromBlock : %d, toBlock : %d ", el.tokenType.String(), fromBlockNum, toBlockNum)
	ethClient := el.ec
	contractAddress := common.HexToAddress(el.contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: fromBlockNum,
		ToBlock:   toBlockNum,
	}

	sub, err := ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		el.errorHandler <- ErrMsg{
			tp:   el.tokenType,
			from: fromBlockNum,
			to:   toBlockNum,
		}
		log.Errorf("erc20 subscribe err : %+v, from : %d, to : %d, type : %s", err, fromBlockNum.Int64(), toBlockNum.Int64(), el.tokenType.String())
		return err
	}
	for _, logEvent := range sub {
		switch logEvent.Topics[0].String() {
		case util.EventSignHash(chain.TRANSFERTOPIC):
			msg := ErrMsg{
				tp:   el.tokenType,
				from: big.NewInt(int64(logEvent.BlockNumber)),
				to:   big.NewInt(int64(logEvent.BlockNumber)),
			}

			input, err := el.abi.Events["Transfer"].Inputs.Unpack(logEvent.Data)
			if err != nil {
				log.Error("erc20 data unpack err : ", err)
				el.errorHandler <- msg
				break
			}
			fromAddr := common.HexToAddress(logEvent.Topics[1].Hex()).String()
			toAddr := common.HexToAddress(logEvent.Topics[2].Hex()).String()
			accept, txType := el.Accept(fromAddr, toAddr)
			if !accept {
				break
			}
			recp, err := el.ec.TransactionReceipt(context.Background(), logEvent.TxHash)
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query txReceipt txHash : %s, err : %+v", logEvent.TxHash, err)
				break
			}
			block, err := el.ec.BlockByNumber(context.Background(), big.NewInt(int64(logEvent.BlockNumber)))
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", logEvent.BlockNumber, err)
				break
			}
			_ = model.SpikeTx{
				From:    fromAddr,
				To:      toAddr,
				TxType:  int64(txType),
				TxHash:  logEvent.TxHash.Hex(),
				Status:  int(recp.Status),
				PayTime: int64(block.Time() * 1000),
				Amount:  input[0].(*big.Int).String(),
			}
		case util.EventSignHash(chain.WITHRAWALTOPIC):
			msg := ErrMsg{
				tp:   el.tokenType,
				from: big.NewInt(int64(logEvent.BlockNumber)),
				to:   big.NewInt(int64(logEvent.BlockNumber)),
			}
			input, err := el.abi.Events["Withdraw"].Inputs.Unpack(logEvent.Data)
			if err != nil {
				log.Error("game vault data unpack err : ", err)
				el.errorHandler <- msg
				break
			}
			if input[0].(common.Address).String() != constant.EmptyAddress {
				break
			}
			fromAddr := input[1].(common.Address).String()
			toAddr := input[2].(common.Address).String()
			accept, txType := el.Accept(fromAddr, toAddr)
			if !accept {
				break
			}
			recp, err := el.ec.TransactionReceipt(context.Background(), logEvent.TxHash)
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query txReceipt txHash : %s, err : %+v", logEvent.TxHash, err)
				break
			}
			block, err := el.ec.BlockByNumber(context.Background(), big.NewInt(int64(logEvent.BlockNumber)))
			if err != nil {
				el.errorHandler <- msg
				log.Errorf("query BlockByNumber blockNum : %d, err : %+v", logEvent.BlockNumber, err)
				break
			}
			_ = model.SpikeTx{
				From:    fromAddr,
				To:      toAddr,
				TxType:  int64(txType),
				TxHash:  logEvent.TxHash.Hex(),
				Status:  int(recp.Status),
				PayTime: int64(block.Time() * 1000),
				Amount:  input[3].(*big.Int).String(),
			}
		}
	}
	return err
}
