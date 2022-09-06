package queryService

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/xerrors"
	"spike-frame/chain/contract"
	"spike-frame/config"
	"spike-frame/response"
	"spike-frame/util"
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
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		erc20Contract, err := contract.NewErc20Contract(common.HexToAddress(config.Cfg.Contract.GovernanceTokenAddress), client)
		if err != nil {
			return
		}
		governanceTokenBalance, err := erc20Contract.BalanceOf(nil, addr)
		if err != nil {
			return
		}
		balanceList = append(balanceList, response.BalanceShow{
			Symbol:  "SKK",
			Balance: util.ParseBalance(governanceTokenBalance),
		})
	}()

	go func() {
		defer wg.Done()
		erc20Contract, err := contract.NewErc20Contract(common.HexToAddress(config.Cfg.Contract.GameTokenAddress), client)
		if err != nil {
			return
		}
		gameTokenBalance, err := erc20Contract.BalanceOf(nil, addr)
		if err != nil {
			return
		}
		balanceList = append(balanceList, response.BalanceShow{
			Symbol:  "SKS",
			Balance: util.ParseBalance(gameTokenBalance),
		})
	}()

	go func() {
		defer wg.Done()
		erc20Contract, err := contract.NewErc20Contract(common.HexToAddress(config.Cfg.Contract.UsdcAddress), client)
		if err != nil {
			return
		}
		usdcBalance, err := erc20Contract.BalanceOf(nil, addr)
		if err != nil {
			return
		}
		balanceList = append(balanceList, response.BalanceShow{
			Symbol:  "USDC",
			Balance: util.ParseBalance(usdcBalance),
		})
	}()

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
	if len(balanceList) != 4 {
		return balanceList, chainNodeError
	}
	log.Infof("wallet : %s  balance : %v", address, balanceList)
	return balanceList, nil
}
