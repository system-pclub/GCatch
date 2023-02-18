package serving6472

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type workItem struct {
	ingressState *ingressState
}

type ingressState struct {
	pendingCount int32
}

type t interface{}

type Interface interface {
	Get() interface{}
	Add(item interface{})
}

type Type struct {
	queue []t
	cond  *sync.Cond
}

func (q *Type) Get() (item interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if len(q.queue) == 0 {
		return nil
	}

	item, q.queue = q.queue[0], q.queue[1:]
	return item
}

func (q *Type) Add(item interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.queue = append(q.queue, item)
}

type DelayingInterface interface {
	Interface
	AddAfter(item interface{})
}

type delayingType struct {
	Interface
}

func (q *delayingType) AddAfter(item interface{}) {
	q.Add(item)
}

func newDelayingQueue() DelayingInterface {
	return &delayingType{&Type{queue: []t{}, cond: sync.NewCond(&sync.Mutex{})}}
}

func NewDelayingQueue() DelayingInterface {
	return newDelayingQueue()
}

type RateLimitingInterface interface {
	DelayingInterface
	AddRateLimited(item interface{})
}

type rateLimitingType struct {
	DelayingInterface
}

func (q *rateLimitingType) AddRateLimited(item interface{}) {
	q.DelayingInterface.AddAfter(item)
}

func NewRateLimitingQueue() RateLimitingInterface {
	return &rateLimitingType{
		DelayingInterface: NewDelayingQueue(),
	}
}

type Prober struct {
	workQueue RateLimitingInterface
}

func (m *Prober) IsReady() {
	workItems := make(map[string][]*workItem)
	ingressState := &ingressState{}
	workItems["0"] = append(workItems["0"], &workItem{
		ingressState: ingressState,
	})
	for _, ipWorkItems := range workItems {
		/*
			go func() {
				m.updateStates(ingressState)
			}()
		*/
		for _, wi := range ipWorkItems {
			m.workQueue.Add(wi)
		}
	}
	ingressState.pendingCount += int32(len(workItems))
}

func (m *Prober) processWorkItem() {
	obj := m.workQueue.Get()
	item, ok := obj.(*workItem)
	if !ok {
		return
	}
	m.updateStates(item.ingressState)
}

func (m *Prober) updateStates(ingressState *ingressState) {
	if atomic.AddInt32(&ingressState.pendingCount, -1) == 0 {
	}
}

func (m *Prober) Start() chan struct{} {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.processWorkItem()
		}()
	}
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func NewProber() *Prober {
	workQueue := NewRateLimitingQueue()
	workQueue.Add(&workItem{&ingressState{}})
	return &Prober{
		workQueue: workQueue,
	}
}

func TestServing6472(t *testing.T) {
	prober := NewProber()
	done := make(chan struct{})
	cancelled := prober.Start()
	defer func() {
		close(done)
		<-cancelled
	}()

	prober.IsReady()
	time.Sleep(1 * time.Millisecond)
}
