package serving3068

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Interface interface {
	Go(func())
	Wait()
}

type impl struct {
	wg     sync.WaitGroup
	workCh chan func()
	once   sync.Once
}

var _ Interface = (*impl)(nil)

func NewWithCapacity(workers, capacity int) Interface {
	i := &impl{
		workCh: make(chan func(), capacity),
	}

	for idx := 0; idx < workers; idx++ {
		go func() {
			for work := range i.workCh {
				func() {
					defer i.wg.Done()
					work()
				}()
			}
		}()
	}

	return i
}

func (i *impl) Go(w func()) {
	i.wg.Add(1)
	i.workCh <- w
}

func (i *impl) Wait() {
	i.once.Do(func() {
		close(i.workCh)

		go func() {
			i.wg.Wait()
		}()
	})
}

func TestServing3068(t *testing.T) {
	p := NewWithCapacity(1, 5)
	wg := &sync.WaitGroup{}
	var cntExecuted int32
	const n = 5
	wg.Add(n)
	go func() {
		for i := 0; i < n; i++ {
			p.Go(func() {
				atomic.AddInt32(&cntExecuted, 1)
			})
			time.Sleep(10 * time.Millisecond)
			wg.Done()
		}
	}()
	p.Wait()
	wg.Wait()
	if cntExecuted == n {
		t.Error("Not all items were expected to execute")
	}
}
