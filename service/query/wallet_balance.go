package query

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spike-engine/spike-web3-server/chain/contract"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/util"
	"golang.org/x/xerrors"
	"sync"
)

var chainNodeError = xerrors.New("chain node error")

type BalanceService struct {
}

var BalanceSrv = new(BalanceService)

func (service *BalanceService) QueryWalletService(address string) ([]response.BalanceShow, error) {
	balanceList := make([]response.BalanceShow, 0)
	client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		return balanceList, chainNodeError
	}
	addr := common.HexToAddress(address)
	contractAddrList := config.Cfg.Contract.ERC20ContractAddress
	var wg sync.WaitGroup
	balanceShowLength := len(contractAddrList) + 1
	wg.Add(balanceShowLength)
	for _, contractAddress := range contractAddrList {
		go func(contractAddr string) {
			defer wg.Done()
			erc20Contract, err := contract.NewErc20Contract(common.HexToAddress(contractAddr), client)
			if err != nil {
				return
			}
			balance, err := erc20Contract.BalanceOf(nil, addr)
			if err != nil {
				return
			}
			symbol, err := erc20Contract.Symbol(nil)
			if err != nil {
				return
			}
			balanceList = append(balanceList, response.BalanceShow{
				Symbol:  symbol,
				Balance: util.ParseBalance(balance),
			})
		}(contractAddress)
	}
	go func() {
		defer wg.Done()
		bnbBalance, err := client.BalanceAt(context.Background(), addr, nil)
		if err != nil {
			return
		}
		balanceList = append(balanceList, response.BalanceShow{
			Symbol:  "BNB",
			Balance: util.ParseBalance(bnbBalance),
		})
	}()
	wg.Wait()
	if len(balanceList) != balanceShowLength {
		return balanceList, chainNodeError
	}
	log.Infof("wallet : %s  balance : %v", address, balanceList)
	return balanceList, nil
}
