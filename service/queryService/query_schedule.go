package queryService

import (
	"context"
	"fmt"
	"golang.org/x/xerrors"
	"spike-frame/model"
	"spike-frame/util"
	"sync"
	"time"
)

var ResourceTable = map[model.TaskType]int{
	model.NftQuery:            1,
	model.Erc20TxRecordQuery:  1,
	model.NativeTxRecordQuery: 2,
}

func (s *Scheduler) getCounter(taskType model.TaskType) (util.Counter, error) {
	switch taskType {
	case model.NftQuery:
		return s.nftListCounter, nil
	case model.NativeTxRecordQuery, model.Erc20TxRecordQuery:
		return s.txRecordCounter, nil
	default:
		return util.Counter{}, xerrors.New(fmt.Sprintf("task type %s is not exist ", taskType.String()))
	}
}

type Scheduler struct {
	schedule        chan *model.QueryRequest
	rq              *model.QueryRequestQueue
	rqLk            sync.Mutex
	nftListCounter  util.Counter
	txRecordCounter util.Counter
}

func NewScheduler() *Scheduler {
	scheduler := &Scheduler{
		schedule:        make(chan *model.QueryRequest),
		rq:              &model.QueryRequestQueue{},
		nftListCounter:  util.NewCounter(12, time.Second),
		txRecordCounter: util.NewCounter(12, time.Second),
	}
	go scheduler.trySched()
	return scheduler
}

func (s *Scheduler) Schedule(ctx context.Context, tp model.TaskType, action model.QueryAction) error {
	done := make(chan model.Result)
	select {
	case s.schedule <- &model.QueryRequest{
		Ctx:    ctx,
		Work:   action,
		Tp:     tp,
		Done:   done,
		Weight: ResourceTable[tp],
	}:
	case <-ctx.Done():
		return ctx.Err()
	}

	select {
	case res := <-done:
		return res.Err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *Scheduler) trySched() {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case req := <-s.schedule:
			s.rqLk.Lock()
			s.rq.Push(req)
			s.rqLk.Unlock()
			log.Infof("---")
		case <-ticker.C:
		}
		s.handleRequest()
	}
}

func (s *Scheduler) handleRequest() {
	s.rqLk.Lock()
	defer s.rqLk.Unlock()
	queueLen := s.rq.Len()
	for i := 0; i < queueLen; i++ {
		task := (*s.rq)[i]
		counter, err := s.getCounter(task.Tp)
		if err != nil {
			log.Errorf("get Counter err :  %v", err.Error())
			continue
		}
		if !counter.Ok(task.Weight) {
			continue
		}
		s.rq.Remove(i)
		go func(t *model.QueryRequest) {
			//do it
			err := t.Work(t.Ctx)
			t.Done <- model.Result{
				Err: err,
			}
		}(task)
	}
}
