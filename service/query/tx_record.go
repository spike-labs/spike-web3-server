package query

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	chain "github.com/spike-engine/spike-web3-server/chain/abi"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/response"
	"github.com/spike-engine/spike-web3-server/util"
	"golang.org/x/xerrors"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"
)

const BscScanRateLimit = "\"Max rate limit reached\""
const txRecordDuration = 10 * time.Minute

func queryNativeTxRecord(address string) (response.BscResult, error) {
	bscRes := response.BscResult{Result: make([]response.TxResult, 0)}

	blockNum, err := util.QueryBlockHeight()
	if err != nil {
		return bscRes, err
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(getNativeUrl(blockNum, address))
	if err != nil {
		return bscRes, err
	}
	err = json.Unmarshal(resp.Body(), &bscRes)
	if err != nil {
		return bscRes, xerrors.New(BscScanRateLimit)
	}
	return bscRes, nil
}

func queryERC20TxRecord(contractAddr, address string) (response.BscResult, error) {
	bscRes := response.BscResult{Result: make([]response.TxResult, 0)}
	blockNum, err := util.QueryBlockHeight()
	if err != nil {
		return bscRes, err
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Accept", "application/json").
		Get(getERC20url(contractAddr, address, blockNum))
	if err != nil {
		return bscRes, err
	}
	err = json.Unmarshal(resp.Body(), &bscRes)

	if err != nil {
		return bscRes, xerrors.New(BscScanRateLimit)
	}
	return bscRes, nil
}

func (qm *QueryManager) handleNativeTxRecordData(walletAddr string, data response.BscResult) (response.BscResult, error) {
	bnbRecord := make([]response.TxResult, 0)
	if len(data.Result) == 0 {
		data.Result = make([]response.TxResult, 0)
		cacheData, _ := json.Marshal(data)
		util.SetFromRedis(walletAddr+constant.NATIVETXRECORDSUFFIX, string(cacheData), txRecordDuration, qm.redisClient)
		return data, nil
	}

	for i := range data.Result {
		if data.Result[i].Input == "0x" {
			bnbRecord = append(bnbRecord, data.Result[i])
			continue
		}
		methodId := data.Result[i].Input[0:10]
		switch methodId {
		case hexutil.Encode(util.GetTxMethodName("swapExactTokensForETHSupportingFeeOnTransferTokens(uint256,uint256,address[],address,uint256)")):
			height, err := strconv.ParseInt(data.Result[i].BlockNumber, 10, 64)
			query := ethereum.FilterQuery{
				FromBlock: big.NewInt(height),
				ToBlock:   big.NewInt(height),
			}
			bscClient, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
			if err != nil {
				log.Errorf("")
				break
			}
			sub, err := bscClient.FilterLogs(context.Background(), query)

			for _, logEvent := range sub {
				if logEvent.Topics[0].String() == util.EventSignHash(chain.WITHRAWALTOPIC) {
					data.Result[i].Type = "Receive"
					value := new(big.Int)
					value.SetString(strings.Split(hexutil.Encode(logEvent.Data), "0x")[1], 16)

					data.Result[i].Value = value.String()
					bnbRecord = append(bnbRecord, data.Result[i])
					break
				}
			}
		case hexutil.Encode(util.GetTxMethodName("swapExactETHForTokens(uint256,address[],address,uint256)")):
			data.Result[i].Type = "Send"
			bnbRecord = append(bnbRecord, data.Result[i])
		}
	}
	sort.Slice(bnbRecord, func(i, j int) bool {
		time1, _ := strconv.Atoi(bnbRecord[i].TimeStamp)
		time2, _ := strconv.Atoi(bnbRecord[j].TimeStamp)
		return time1 > time2
	})
	data.Result = bnbRecord
	cacheData, _ := json.Marshal(data)
	util.SetFromRedis(walletAddr+constant.NATIVETXRECORDSUFFIX, string(cacheData), txRecordDuration, qm.redisClient)
	return data, nil
}

func (qm *QueryManager) handleERC20TxRecordData(walletAddr string, contractAddr string, data response.BscResult) (response.BscResult, error) {
	if len(data.Result) == 0 {
		data.Result = make([]response.TxResult, 0)
		cacheData, _ := json.Marshal(data)
		util.SetFromRedis(walletAddr+contractAddr+constant.ERC20TXRECORDSUFFIX, string(cacheData), txRecordDuration, qm.redisClient)
		return data, nil
	}
	cacheData, _ := json.Marshal(data)
	util.SetFromRedis(walletAddr+contractAddr+constant.ERC20TXRECORDSUFFIX, string(cacheData), txRecordDuration, qm.redisClient)
	return data, nil
}

func getNativeUrl(blockNumber uint64, address string) string {
	return fmt.Sprintf("%s?module=account&action=txlist&address=%s&startblock=%d&endblock=%d&offset=10000&page=1&sort=desc&apikey=%s", config.Cfg.BscScan.UrlPrefix, address, blockNumber-201600, blockNumber, config.Cfg.BscScan.ApiKey)
}

func getNativeInternalUrl(blockNumber uint64, address string) string {
	return fmt.Sprintf("%s?module=account&action=txlistinternal&address=%s&startblock=%d&endblock=%d&offset=10000&page=1&sort=desc&apikey=%s", config.Cfg.BscScan.UrlPrefix, address, blockNumber-201600, blockNumber, config.Cfg.BscScan.ApiKey)
}

func getERC20url(contractAddr, addr string, blockNumber uint64) string {
	return fmt.Sprintf("%s?module=account&action=tokentx&address=%s&startblock=%d&endblock=%d&offset=10000&page=1&sort=desc&apikey=%s&contractaddress=%s", config.Cfg.BscScan.UrlPrefix, addr, blockNumber-201600, blockNumber, config.Cfg.BscScan.ApiKey, contractAddr)
}
