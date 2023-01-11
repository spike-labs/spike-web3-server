package model

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type BatchMintReq struct {
	Uuid       string `json:"uuid"`
	TokenID    int64  `json:"token_id"`
	TokenURI   string `json:"token_uri"`
	NFTAddress string `json:"nft_address"`
}

type BatchMintQueue struct {
	Reqs map[string][]BatchMintReq
	qLK  sync.Mutex
}

func (q *BatchMintQueue) Push(x BatchMintReq) {
	q.qLK.Lock()
	defer q.qLK.Unlock()
	item := x
	q.Reqs[x.NFTAddress] = append(q.Reqs[x.NFTAddress], item)
}

func (q *BatchMintQueue) Remove(nftAddress string, i int) BatchMintReq {
	q.qLK.Lock()
	defer q.qLK.Unlock()
	old := q.Reqs[nftAddress]
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = BatchMintReq{}
	q.Reqs[nftAddress] = old[0 : n-1]
	return item
}

func (q *BatchMintQueue) Clear(nftAddress ...string) {
	q.qLK.Lock()
	defer q.qLK.Unlock()

	for i := range nftAddress {
		q.Reqs[nftAddress[i]] = make([]BatchMintReq, 0)
	}
}

func (q *BatchMintQueue) Len(nftAddress string) int {
	return len(q.Reqs[nftAddress])
}

func (q *BatchMintQueue) LenOfAll() int {
	var taskLen int
	for i := range q.Reqs {
		taskLen += len(q.Reqs[i])
	}

	return taskLen
}

type WithdrawTokenReq struct {
	Uuid         string         `json:"uuid"`
	ToAddress    common.Address `json:"to_address"`
	Amount       string         `json:"amount"`
	TokenAddress common.Address `json:"token_address"`
}
type WithdrawTokenQueue struct {
	wLK  sync.Mutex
	Reqs []WithdrawTokenReq
}

func (q *WithdrawTokenQueue) Push(x WithdrawTokenReq) {
	item := x
	q.Reqs = append(q.Reqs, item)
}

func (q *WithdrawTokenQueue) Remove(i int) WithdrawTokenReq {
	old := q.Reqs
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = WithdrawTokenReq{}
	q.Reqs = old[0 : n-1]
	return item
}
func (q *WithdrawTokenQueue) Clear() {
	q.Reqs = make([]WithdrawTokenReq, 0)
}

func (q *WithdrawTokenQueue) Len() int { return len(q.Reqs) }

type WithdrawNFTReq struct {
	Uuid         string         `json:"uuid"`
	TokenId      int64          `json:"token_id"`
	ToAddress    common.Address `json:"to_address"`
	TokenAddress common.Address `json:"token_address"`
}
type WithdrawNFTQueue struct {
	wLK  sync.Mutex
	Reqs []WithdrawNFTReq
}

func (q *WithdrawNFTQueue) Push(x WithdrawNFTReq) {
	item := x
	q.Reqs = append(q.Reqs, item)
}

func (q *WithdrawNFTQueue) Remove(i int) WithdrawNFTReq {
	old := q.Reqs
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = WithdrawNFTReq{}
	q.Reqs = old[0 : n-1]
	return item
}
func (q *WithdrawNFTQueue) Clear() {
	q.Reqs = make([]WithdrawNFTReq, 0)
}

func (q *WithdrawNFTQueue) Len() int { return len(q.Reqs) }
