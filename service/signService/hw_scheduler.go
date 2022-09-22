package signService

import (
	"sort"
	"github.com/spike-engine/spike-web3-server/config"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/game"
	"github.com/spike-engine/spike-web3-server/global"
	"github.com/spike-engine/spike-web3-server/model"
	"sync"
	"time"
)

type hotWalletScheduler struct {
	gorm game.TxTracker

	taskLK  sync.RWMutex
	wLK     sync.Mutex
	workers []Worker

	mintSched  chan *model.BatchMintReq
	tokenSched chan *model.WithdrawTokenReq
	nftSched   chan *model.WithdrawNFTReq

	mintQueue  *model.BatchMintQueue
	tokenQueue *model.WithdrawTokenQueue
	nftQueue   *model.WithdrawNFTQueue
}

func newScheduler() *hotWalletScheduler {
	return &hotWalletScheduler{
		gorm:    global.DbAccessor,
		workers: make([]Worker, 0),

		mintSched: make(chan *model.BatchMintReq),
		mintQueue: &model.BatchMintQueue{Reqs: make([]model.BatchMintReq, 0)},

		tokenSched: make(chan *model.WithdrawTokenReq),
		tokenQueue: &model.WithdrawTokenQueue{Reqs: make([]model.WithdrawTokenReq, 0)},

		nftSched: make(chan *model.WithdrawNFTReq),
		nftQueue: &model.WithdrawNFTQueue{Reqs: make([]model.WithdrawNFTReq, 0)},
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
			hw.mintQueue.Push(*req)
			if hw.mintQueue.Len() < config.Cfg.SignService.TaskThreshold {
				break
			}
			hw.schedMintTask()

		case req := <-hw.tokenSched:
			hw.tokenQueue.Push(*req)
			if hw.tokenQueue.Len() < config.Cfg.SignService.TaskThreshold {
				break
			}
			hw.schedTokenTask()

		case req := <-hw.nftSched:
			hw.nftQueue.Push(*req)
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
		ticker.Reset(time.Duration(config.Cfg.SignService.SchedInterval) * time.Minute)

	}

}

func (hw *hotWalletScheduler) schedMintTask() {
	for _, reqs := range hw.CheckExecMintTask() {
		go func(q []model.BatchMintReq) {
			var txStatus int
			uuids, TxHash, err := hw.pickRightWorker().BatchMint(q)
			if err != nil {
				txStatus = constant.TXFAILED
				log.Error("===Spike log:filed exec batchMint. error: ", err)
				return
			}
			txStatus = constant.ORDERHANDLED
			err = hw.gorm.RecordTxHash(uuids, TxHash, txStatus)
			if err != nil {
				log.Error("===Spike log:filed exec batchMint. error: ", err)
				return
			}
			log.Info("===Spike log: success exec batchMint")
		}(reqs)
	}
}

func (hw *hotWalletScheduler) schedTokenTask() {
	for _, queue := range hw.CheckExecTokenTask() {
		go func(q []model.WithdrawTokenReq) {
			var txStatus int
			uuids, TxHash, err := hw.pickRightWorker().WithdrawToken(q)
			if err != nil {
				txStatus = constant.TXFAILED
				log.Error("===Spike log:filed exec batchWithdrawToken. error :", err)
				return
			}
			txStatus = constant.ORDERHANDLED
			err = hw.gorm.RecordTxHash(uuids, TxHash, txStatus)
			if err != nil {
				log.Error("===Spike log:filed exec batchWithdrawToken. error :", err)
				return
			}
			log.Info("===Spike log: success exec batchWithdrawToken")
		}(queue)
	}
}

func (hw *hotWalletScheduler) schedNFTTask() {
	for _, queue := range hw.CheckExecNFTTask() {
		go func(q []model.WithdrawNFTReq) {
			var txStatus int
			uuids, TxHash, err := hw.pickRightWorker().WithdrawNFT(q)
			if err != nil {
				txStatus = constant.TXFAILED
				log.Error("===Spike log:filed exec batchWithdrawNFT. error: ", err)
				return
			}
			txStatus = constant.ORDERHANDLED
			err = hw.gorm.RecordTxHash(uuids, TxHash, txStatus)
			if err != nil {
				log.Error("===Spike log:filed exec batchWithdrawNFT. error: ", err)
				return
			}
			log.Info("===Spike log: success exec batchWithdrawNFT")
		}(queue)
	}
}

func (hw *hotWalletScheduler) pickRightWorker() Worker {
	hw.wLK.Lock()
	defer hw.wLK.Unlock()

	sort.Slice(hw.workers, func(i, j int) bool {
		return hw.workers[i].GetInfo().TaskNum < hw.workers[j].GetInfo().TaskNum
	})

	hw.workers[0].AddTaskNum()
	return hw.workers[0]
}

func (hw *hotWalletScheduler) CheckExecMintTask() [][]model.BatchMintReq {
	hw.taskLK.Lock()
	defer hw.taskLK.Unlock()
	taskNum := hw.mintQueue.Len()
	taskQueues := make([][]model.BatchMintReq, 0)
	switch {

	case taskNum > config.Cfg.SignService.TaskThreshold:

		for i := 0; i < taskNum/config.Cfg.SignService.TaskThreshold; i++ {

			reqs := hw.mintQueue.Reqs[i*config.Cfg.SignService.TaskThreshold : (i+1)*config.Cfg.SignService.TaskThreshold]
			taskQueues = append(taskQueues, reqs)

			if i+1 == taskNum/config.Cfg.SignService.TaskThreshold && taskNum%config.Cfg.SignService.TaskThreshold != 0 {
				reqs := hw.mintQueue.Reqs[(i+1)*config.Cfg.SignService.TaskThreshold:]
				taskQueues = append(taskQueues, reqs)
			}
		}
		hw.mintQueue.Clear()
		return taskQueues
	case taskNum <= config.Cfg.SignService.TaskThreshold:
		taskQueues = append(taskQueues, hw.mintQueue.Reqs)
		hw.mintQueue.Clear()
		return taskQueues
	}
	return nil
}

func (hw *hotWalletScheduler) CheckExecTokenTask() [][]model.WithdrawTokenReq {
	hw.taskLK.Lock()
	defer hw.taskLK.Unlock()

	taskNum := hw.tokenQueue.Len()
	taskQueues := make([][]model.WithdrawTokenReq, 0)
	switch {
	case taskNum > config.Cfg.SignService.TaskThreshold:

		for i := 0; i < taskNum/config.Cfg.SignService.TaskThreshold; i++ {

			reqs := hw.tokenQueue.Reqs[i*config.Cfg.SignService.TaskThreshold : (i+1)*config.Cfg.SignService.TaskThreshold]
			taskQueues = append(taskQueues, reqs)

			if i+1 == taskNum/config.Cfg.SignService.TaskThreshold && taskNum%config.Cfg.SignService.TaskThreshold != 0 {
				reqs := hw.tokenQueue.Reqs[(i+1)*config.Cfg.SignService.TaskThreshold:]
				taskQueues = append(taskQueues, reqs)
			}
		}
		hw.tokenQueue.Clear()
		return taskQueues
	case taskNum <= config.Cfg.SignService.TaskThreshold:
		taskQueues = append(taskQueues, hw.tokenQueue.Reqs)
		hw.tokenQueue.Clear()
		return taskQueues
	}
	return nil
}

func (hw *hotWalletScheduler) CheckExecNFTTask() [][]model.WithdrawNFTReq {
	hw.taskLK.Lock()
	defer hw.taskLK.Unlock()

	taskNum := hw.nftQueue.Len()
	taskQueues := make([][]model.WithdrawNFTReq, 0)
	switch {
	case taskNum > config.Cfg.SignService.TaskThreshold:

		for i := 0; i < taskNum/config.Cfg.SignService.TaskThreshold; i++ {

			reqs := hw.nftQueue.Reqs[i*config.Cfg.SignService.TaskThreshold : (i+1)*config.Cfg.SignService.TaskThreshold]
			taskQueues = append(taskQueues, reqs)

			if i+1 == taskNum/config.Cfg.SignService.TaskThreshold && taskNum%config.Cfg.SignService.TaskThreshold != 0 {
				reqs := hw.nftQueue.Reqs[(i+1)*config.Cfg.SignService.TaskThreshold:]
				taskQueues = append(taskQueues, reqs)
			}
		}
		hw.nftQueue.Clear()
		return taskQueues
	case taskNum <= config.Cfg.SignService.TaskThreshold:
		taskQueues = append(taskQueues, hw.nftQueue.Reqs)
		hw.nftQueue.Clear()
		return taskQueues
	}
	return nil
}
