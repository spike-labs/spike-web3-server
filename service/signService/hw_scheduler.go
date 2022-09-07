package signService

import (
	"sort"
	"spike-frame/config"
	"spike-frame/constant"
	"spike-frame/dao"
	"spike-frame/model"
	"sync"
	"time"
)

type hotWalletScheduler struct {
	gorm *dao.GormAccessor

	workerLK sync.RWMutex
	workers  []Worker

	mintSched  chan *model.BatchMintReq
	tokenSched chan *model.WithdrawTokenReq
	nftSched   chan *model.WithdrawNFTReq

	mintQueue  *model.BatchMintQueue
	tokenQueue *model.WithdrawTokenQueue
	nftQueue   *model.WithdrawNFTQueue
}

func newScheduler() *hotWalletScheduler {
	return &hotWalletScheduler{
		gorm:    dao.NewGormAccessor(constant.GormClient),
		workers: make([]Worker, 0),

		mintSched: make(chan *model.BatchMintReq),
		mintQueue: &model.BatchMintQueue{Reqs: make([]*model.BatchMintReq, 0)},

		tokenSched: make(chan *model.WithdrawTokenReq),
		tokenQueue: &model.WithdrawTokenQueue{Reqs: make([]*model.WithdrawTokenReq, 0)},

		nftSched: make(chan *model.WithdrawNFTReq),
		nftQueue: &model.WithdrawNFTQueue{Reqs: make([]*model.WithdrawNFTReq, 0)},
	}
}

func (hw *hotWalletScheduler) Schedule(req interface{}) {

	switch req.(type) {
	case *model.BatchMintReq:
		mintReq := req.(*model.BatchMintReq)
		hw.mintSched <- mintReq
	case *model.WithdrawTokenReq:
		tokenReq := req.(*model.WithdrawTokenReq)
		hw.tokenSched <- tokenReq
	case *model.WithdrawNFTReq:
		nftReq := req.(*model.WithdrawNFTReq)
		hw.nftSched <- nftReq
	}
}

func (hw *hotWalletScheduler) runSchedule() {

	ticker := time.NewTicker(time.Duration(config.Cfg.SignService.SchedInterval) * time.Minute)
	for {
		select {
		case req := <-hw.mintSched:
			hw.mintQueue.Push(req)
			if hw.mintQueue.Len() < config.Cfg.SignService.TaskThreshold {
				break
			}
			hw.schedMintTask()

		case req := <-hw.tokenSched:
			hw.tokenQueue.Push(req)
			if hw.tokenQueue.Len() < config.Cfg.SignService.TaskThreshold {
				break
			}
			hw.schedTokenTask()

		case req := <-hw.nftSched:
			hw.nftQueue.Push(req)
			if hw.nftQueue.Len() < config.Cfg.SignService.TaskThreshold {
				break
			}
			hw.schedNFTTask()

		case <-ticker.C:
			if hw.mintQueue.Len() != 0 {
				hw.schedMintTask()
			}

			if hw.tokenQueue.Len() != 0 {
				hw.schedTokenTask()
			}

			if hw.nftQueue.Len() != 0 {
				hw.schedNFTTask()
			}
		}

	}

}

func (hw *hotWalletScheduler) schedMintTask() {
	for _, queue := range hw.mintQueue.CheckExecTask() {
		go func(q *model.BatchMintQueue) {
			var txStatus int
			uuids, TxHash, err := hw.pickRightWorker().BatchMint(q)
			if err != nil {
				txStatus = constant.TXFAILED
				log.Error("===Spike log:", err)
				return
			}
			txStatus = constant.ORDERHANDLED
			err = hw.gorm.RecordTxHash(uuids, TxHash, txStatus)
			if err != nil {
				log.Error("===Spike log:", err)
				return
			}
		}(queue)
	}
}

func (hw *hotWalletScheduler) schedTokenTask() {
	for _, queue := range hw.tokenQueue.CheckExecTask() {
		go func(q *model.WithdrawTokenQueue) {
			var txStatus int
			uuids, TxHash, err := hw.pickRightWorker().WithdrawToken(q)
			if err != nil {
				txStatus = constant.TXFAILED
				log.Error("===Spike log:", err)
				return
			}
			txStatus = constant.ORDERHANDLED
			err = hw.gorm.RecordTxHash(uuids, TxHash, txStatus)
			if err != nil {
				log.Error("===Spike log:", err)
				return
			}
		}(queue)
	}
}

func (hw *hotWalletScheduler) schedNFTTask() {
	for _, queue := range hw.nftQueue.CheckExecTask() {
		go func(q *model.WithdrawNFTQueue) {
			var txStatus int
			uuids, TxHash, err := hw.pickRightWorker().WithdrawNFT(q)
			if err != nil {
				txStatus = constant.TXFAILED
				log.Error("===Spike log:", err)
				return
			}
			txStatus = constant.ORDERHANDLED
			err = hw.gorm.RecordTxHash(uuids, TxHash, txStatus)
			if err != nil {
				log.Error("===Spike log:", err)
				return
			}
		}(queue)
	}
}

func (hw *hotWalletScheduler) pickRightWorker() Worker {
	sort.Slice(hw.workers, func(i, j int) bool {
		return hw.workers[i].GetInfo().TaskNum < hw.workers[j].GetInfo().TaskNum
	})

	return hw.workers[0]
}
