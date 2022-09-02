package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"sync"

	"spike-frame/config"
)

type BatchMintReq struct {
	Uuid     uuid.UUID `json:"uuid"`
	TokenID  int64     `json:"token_id"`
	TokenURI string    `json:"token_uri"`
}

type BatchMintQueue struct {
	Reqs []*BatchMintReq
	qLK  sync.Mutex
}

func (q *BatchMintQueue) Push(x *BatchMintReq) {
	q.qLK.Lock()
	defer q.qLK.Unlock()
	item := x
	q.Reqs = append(q.Reqs, item)
}

func (q *BatchMintQueue) Remove(i int) *BatchMintReq {
	q.qLK.Lock()
	defer q.qLK.Unlock()
	old := q.Reqs
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	q.Reqs = old[0 : n-1]
	return item
}
func (q *BatchMintQueue) Clear() {
	q.qLK.Lock()
	defer q.qLK.Unlock()
	q = &BatchMintQueue{}
}

func (q *BatchMintQueue) Len() int { return len(q.Reqs) }

func (q *BatchMintQueue) CheckExecTask() []*BatchMintQueue {
	taskNum := q.Len()
	taskQueues := make([]*BatchMintQueue, 0)
	switch {

	case taskNum > config.Cfg.SignService.TaskThreshold:

		for i := 0; i < taskNum/config.Cfg.SignService.TaskThreshold; i++ {

			reqs := q.Reqs[i*config.Cfg.SignService.TaskThreshold : (i+1)*config.Cfg.SignService.TaskThreshold]
			taskQueues = append(taskQueues, &BatchMintQueue{Reqs: reqs})

			if i+1 == taskNum/config.Cfg.SignService.TaskThreshold && taskNum%config.Cfg.SignService.TaskThreshold != 0 {
				reqs := q.Reqs[(i+1)*config.Cfg.SignService.TaskThreshold:]
				taskQueues = append(taskQueues, &BatchMintQueue{Reqs: reqs})
			}
		}
		q.Clear()
		return taskQueues
	case taskNum <= config.Cfg.SignService.TaskThreshold:
		taskQueues = append(taskQueues, q)
		q.Clear()
		return taskQueues
	}
	return nil
}

type WithdrawTokenReq struct {
	Uuid         uuid.UUID      `json:"uuid"`
	ToAddress    common.Address `json:"to_address"`
	Amount       int64          `json:"amount"`
	TokenAddress common.Address `json:"token_address"`
}
type WithdrawTokenQueue struct {
	wLK  sync.Mutex
	Reqs []*WithdrawTokenReq
}

func (q *WithdrawTokenQueue) Push(x *WithdrawTokenReq) {
	item := x
	q.Reqs = append(q.Reqs, item)
}

func (q *WithdrawTokenQueue) Remove(i int) *WithdrawTokenReq {
	old := q.Reqs
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	q.Reqs = old[0 : n-1]
	return item
}
func (q *WithdrawTokenQueue) Clear() {
	q = &WithdrawTokenQueue{}
}

func (q *WithdrawTokenQueue) Len() int { return len(q.Reqs) }

func (q *WithdrawTokenQueue) CheckExecTask() []*WithdrawTokenQueue {
	taskNum := q.Len()
	taskQueues := make([]*WithdrawTokenQueue, 0)
	switch {
	case taskNum > config.Cfg.SignService.TaskThreshold:

		for i := 0; i < taskNum/config.Cfg.SignService.TaskThreshold; i++ {

			reqs := q.Reqs[i*config.Cfg.SignService.TaskThreshold : (i+1)*config.Cfg.SignService.TaskThreshold]
			taskQueues = append(taskQueues, &WithdrawTokenQueue{Reqs: reqs})

			if i+1 == taskNum/config.Cfg.SignService.TaskThreshold && taskNum%config.Cfg.SignService.TaskThreshold != 0 {
				reqs := q.Reqs[(i+1)*config.Cfg.SignService.TaskThreshold:]
				taskQueues = append(taskQueues, &WithdrawTokenQueue{Reqs: reqs})
			}
		}
		q.Clear()
		return taskQueues
	case taskNum <= config.Cfg.SignService.TaskThreshold:
		taskQueues = append(taskQueues, q)
		q.Clear()
		return taskQueues
	}
	return nil
}

type WithdrawNFTReq struct {
	Uuid         uuid.UUID      `json:"uuid"`
	TokenId      int64          `json:"token_id"`
	ToAddress    common.Address `json:"to_address"`
	TokenAddress common.Address `json:"token_address"`
}
type WithdrawNFTQueue struct {
	wLK  sync.Mutex
	Reqs []*WithdrawNFTReq
}

func (q *WithdrawNFTQueue) Push(x *WithdrawNFTReq) {
	item := x
	q.Reqs = append(q.Reqs, item)
}

func (q *WithdrawNFTQueue) Remove(i int) *WithdrawNFTReq {
	old := q.Reqs
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	q.Reqs = old[0 : n-1]
	return item
}
func (q *WithdrawNFTQueue) Clear() {
	q = &WithdrawNFTQueue{}
}

func (q *WithdrawNFTQueue) Len() int { return len(q.Reqs) }

func (q *WithdrawNFTQueue) CheckExecTask() []*WithdrawNFTQueue {
	taskNum := q.Len()
	taskQueues := make([]*WithdrawNFTQueue, 0)
	switch {
	case taskNum > config.Cfg.SignService.TaskThreshold:

		for i := 0; i < taskNum/config.Cfg.SignService.TaskThreshold; i++ {

			reqs := q.Reqs[i*config.Cfg.SignService.TaskThreshold : (i+1)*config.Cfg.SignService.TaskThreshold]
			taskQueues = append(taskQueues, &WithdrawNFTQueue{Reqs: reqs})

			if i+1 == taskNum/config.Cfg.SignService.TaskThreshold && taskNum%config.Cfg.SignService.TaskThreshold != 0 {
				reqs := q.Reqs[(i+1)*config.Cfg.SignService.TaskThreshold:]
				taskQueues = append(taskQueues, &WithdrawNFTQueue{Reqs: reqs})
			}
		}
		q.Clear()
		return taskQueues
	case taskNum <= config.Cfg.SignService.TaskThreshold:
		taskQueues = append(taskQueues, q)
		q.Clear()
		return taskQueues
	}
	return nil
}
