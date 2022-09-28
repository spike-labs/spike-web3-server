package game

import (
	"github.com/google/uuid"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/global"
	"github.com/spike-engine/spike-web3-server/model"
	"github.com/spike-engine/spike-web3-server/util"
	"strconv"
	"sync"
	"time"
)

var (
	nmMgr     *NftManager
	nmMgrOnce sync.Once
)

type NftManager struct {
	NftOwnerTracker
}

type NftOwnerUpdateEvent struct {
	OwnerAddr    string
	ContractAddr string
	TokenId      int64
	UpdateTime   int64
}

func NewNftManager(tracker NftOwnerTracker) *NftManager {
	nmMgrOnce.Do(func() {
		nmMgr = &NftManager{tracker}
	})
	return nmMgr
}

func (nm *NftManager) Update(event interface{}) {

	e, ok := event.(NftOwnerUpdateEvent)
	log.Infof("e :%v", e)
	if !ok {
		return
	}
	util.Lock(strconv.FormatInt(e.TokenId, 10), constant.TXCBVALUE, LOCKTIMEOUTDURATION, global.RedisClient)

	defer util.UnLock(strconv.FormatInt(e.TokenId, 10), global.RedisClient)

	nftOwner, err := nm.QueryNftOwner(e.TokenId, e.ContractAddr)
	if err != nil {
		log.Errorf("query nft owner err : %v, event : %v", err, e)
		return
	}
	if len(nftOwner) == 0 {
		if err := nm.AddNftOwner(model.NftOwner{
			Id:              uuid.New().String(),
			OwnerAddress:    e.OwnerAddr,
			ContractAddress: e.ContractAddr,
			UpdateTime:      e.UpdateTime,
			TokenId:         e.TokenId,
		}); err != nil {
			log.Errorf("add nft owner err, tokenId : %d , contractAddr : %s, ownerAddr : %s", e.TokenId, e.ContractAddr, e.OwnerAddr)
		}
		return
	}
	if time.UnixMilli(nftOwner[0].UpdateTime).After(time.UnixMilli(e.UpdateTime)) {
		log.Errorf("current event is outdated, event : %v", e)
		return
	}
	if nftOwner[0].OwnerAddress == e.OwnerAddr {
		log.Infof("nft owner has been updated, tokenId : %d , contractAddr : %s, ownerAddr : %s", e.TokenId, e.ContractAddr, e.OwnerAddr)
		return
	}
	if err := nm.UpdateNftOwner(e.OwnerAddr, e.ContractAddr, e.TokenId, e.UpdateTime); err != nil {
		log.Errorf("update nft owner err, tokenId : %d , contractAddr : %s, ownerAddr : %s", e.TokenId, e.ContractAddr, e.OwnerAddr)
		return
	}
}
