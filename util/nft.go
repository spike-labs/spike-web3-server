package util

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-resty/resty/v2"
	"math/big"
	"spike-frame/chain/contract"
	"spike-frame/config"
	"spike-frame/model"
	"spike-frame/response"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ConvertNftResult(gameNftAddress string, res []response.NftResult) []response.NftResult {
	client, err := ethclient.Dial(config.Cfg.Chain.RpcNodeAddress)
	if err != nil {
		return res
	}
	gameNft, err := contract.NewErc721Contract(common.HexToAddress(gameNftAddress), client)
	if err != nil {
		log.Error("new auNft err : ", err)
		return res
	}
	throttle := make(chan struct{}, 20)
	var wg sync.WaitGroup
	for i, nftResult := range res {
		wg.Add(1)
		throttle <- struct{}{}
		go func(k int, v response.NftResult) {
			defer func() {
				wg.Done()
				<-throttle
			}()

			if v.TokenUri == "" || v.Metadata == "" {
				tokenId, err := strconv.Atoi(v.TokenId)
				if err != nil {
					log.Errorf("string %s convert int err : %v", v.TokenId, err)
					return
				}
				uri, err := gameNft.TokenURI(nil, big.NewInt(int64(tokenId)))
				if err != nil {
					log.Errorf("query tokenUri tokenId : %d, err : %+v", tokenId, err)
					return
				}
				client := resty.New()
				t3 := time.Now()
				resp, err := client.R().Get(uri)
				if err != nil {
					log.Errorf("req err : uri :%s", uri)
					return
				}
				log.Info("query metadata took :", time.Since(t3))
				log.Infof("query nft tokenId : %d, uri : %s", tokenId, uri)
				var m model.Metadata
				err = json.Unmarshal(resp.Body(), &m)
				if err != nil {
					log.Errorf("json unmarshal err : %+v, resp : %s", err, string(resp.Body()))
					return
				}
				metadata, err := json.Marshal(m)
				if err != nil {
					log.Errorf("json marshal err : %+v, meatadata : %+v", err, m)
					return
				}
				res[k].TokenUri = uri
				res[k].Metadata = string(metadata)
			}
		}(i, nftResult)
	}
	wg.Wait()
	return res
}

func ParseCacheData(cds []model.CacheData) map[string][]model.CacheData {
	dataMap := make(map[string][]model.CacheData)
	for _, v := range cds {
		if _, have := dataMap[v.Type]; have {
			dataMap[v.Type] = append(dataMap[v.Type], v)
		} else {
			var cd []model.CacheData
			cd = append(cd, v)
			dataMap[v.Type] = cd
		}
	}
	return dataMap
}

func ParseMetadata(nr []response.NftResult) []model.CacheData {
	var dataList []model.CacheData
	for _, v := range nr {
		var cd model.CacheData
		cd.TokenId = v.TokenId
		cd.BlockNumber = v.BlockNumber
		var m model.Metadata
		err := json.Unmarshal([]byte(v.Metadata), &m)
		if err != nil {
			log.Error("json unmarshal err : ", err, v.TokenId, v.TokenUri)
			continue
		}
		split := strings.Split(m.Name, " #")
		if len(split) != 2 {
			log.Errorf("nft name len != 2")
			continue
		}
		cd.Type = split[0]
		cd.GameId = split[1]
		cd.Image = m.Image
		cd.Description = m.Description
		cd.SpikeInfo = m.SpikeInfo
		attrMap := make(map[string]interface{})
		for _, attr := range m.Attribute {
			attrMap[attr.TraitType] = attr.Value
		}
		cd.Attributes = attrMap
		dataList = append(dataList, cd)
	}
	return dataList
}
