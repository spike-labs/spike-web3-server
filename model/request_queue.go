package model

import "context"

type QueryAction func(ctx context.Context) error

type Result struct {
	R   interface{}
	Err error
}

type QueryRequest struct {
	Ctx    context.Context
	Work   QueryAction
	Tp     TaskType
	Done   chan Result
	Weight int
}

type QueryRequestQueue []*QueryRequest

func (q *QueryRequestQueue) Push(x *QueryRequest) {
	item := x
	*q = append(*q, item)
}

func (q *QueryRequestQueue) Remove(i int) *QueryRequest {
	old := *q
	n := len(old)
	item := old[i]
	old[i] = old[n-1]
	old[n-1] = nil
	*q = old[0 : n-1]
	return item
}

func (q QueryRequestQueue) Len() int { return len(q) }
