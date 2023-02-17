package kubernetes81148

import (
	"sync"
	"testing"
	"time"
)

const unschedulableQTimeInterval = 60 * time.Second

type Pod string

type PodInfo struct {
	Pod       Pod
	Timestamp time.Time
}

type UnschedulablePodsMap struct {
	podInfoMap map[string]*PodInfo
	keyFunc    func(Pod) string
}

func (u *UnschedulablePodsMap) addOrUpdate(pInfo *PodInfo) {
	podID := u.keyFunc(pInfo.Pod)
	u.podInfoMap[podID] = pInfo
}

func GetPodFullName(pod Pod) string {
	return string(pod)
}

func newUnschedulablePodsMap() *UnschedulablePodsMap {
	return &UnschedulablePodsMap{
		podInfoMap: make(map[string]*PodInfo),
		keyFunc:    GetPodFullName,
	}
}

type PriorityQueue struct {
	stop           <-chan struct{}
	lock           sync.RWMutex
	unschedulableQ *UnschedulablePodsMap
}

func (p *PriorityQueue) flushUnschedulableQLeftover() {
	p.lock.Lock()
	defer p.lock.Unlock()

	for _, pInfo := range p.unschedulableQ.podInfoMap {
		_ = pInfo.Timestamp
	}
}

func (p *PriorityQueue) run() {
	go Until(p.flushUnschedulableQLeftover, p.stop)
}

func (p *PriorityQueue) newPodInfo(pod Pod) *PodInfo {
	return &PodInfo{
		Pod:       pod,
		Timestamp: time.Now(),
	}
}

func NewPriorityQueueWithClock(stop <-chan struct{}) *PriorityQueue {
	pq := &PriorityQueue{
		stop:           stop,
		unschedulableQ: newUnschedulablePodsMap(),
	}
	pq.run()
	return pq
}

func NewPriorityQueue(stop <-chan struct{}) *PriorityQueue {
	return NewPriorityQueueWithClock(stop)
}

func BackoffUntil(f func(), stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
		}

		func() {
			f()
		}()

		select {
		case <-stopCh:
			return
		}
	}
}

func JitterUntil(f func(), stopCh <-chan struct{}) {
	BackoffUntil(f, stopCh)
}

func Until(f func(), stopCh <-chan struct{}) {
	JitterUntil(f, stopCh)
}

func addOrUpdateUnschedulablePod(p *PriorityQueue, pod Pod) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.unschedulableQ.addOrUpdate(p.newPodInfo(pod))
}

func TestKubernetes81148(t *testing.T) {
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		q := NewPriorityQueue(stop)
		highPod := Pod("1")
		addOrUpdateUnschedulablePod(q, highPod)
		q.unschedulableQ.podInfoMap[GetPodFullName(highPod)].Timestamp = time.Now().Add(-1 * unschedulableQTimeInterval)
	}()
	wg.Wait()
	close(stop)
}
