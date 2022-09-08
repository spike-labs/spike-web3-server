package util

import (
	"sync"
	"time"
)

type Counter struct {
	rate  int
	begin time.Time
	cycle time.Duration
	count int
	lock  sync.Mutex
}

func (l *Counter) Ok(weight, rate int) bool {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.count > l.rate-1 {
		now := time.Now()
		if now.Sub(l.begin) >= l.cycle {
			l.Reset(now, rate)
			return true
		} else {
			log.Infof("rate limit reached")
			return false
		}
	} else {
		l.count = l.count + weight
		return true
	}
}

func NewCounter(r int, cycle time.Duration) Counter {
	return Counter{
		rate:  r,
		begin: time.Now(),
		cycle: cycle,
		count: 0,
	}
}

func (l *Counter) Reset(t time.Time, rate int) {
	l.begin = t
	l.rate = rate
	l.count = 0
}
