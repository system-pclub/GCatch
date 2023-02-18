/*
 * Project: kubernetes
 * Tag: Reproduce misbehavior
 * Issue or PR  : https://github.com/kubernetes/kubernetes/pull/58107
 * Buggy version: 2f17d782eb2772d6401da7ddced9ac90656a7a79
 * fix commit-id: 010a127314a935d8d038f8dd4559fc5b249813e4
 * Flaky: 53/100
 * Description:
 *   The rules for read and write lock: allows concurrent read lock;
 * write lock has higher priority than read lock.
 *   There are two queues (queue 1 and queue 2) involved in this bug,
 * and the two queues are protected by the same read-write lock
 * (rq.workerLock.RLock()). Before getting an element from queue 1 or
 * queue 2, rq.workerLock.RLock() is acquired. If the queue is empty,
 * cond.Wait() will be invoked. There is another goroutine (goroutine D),
 * which will periodically invoke rq.workerLock.Lock(). Under the following
 * situation, deadlock will happen. Queue 1 is empty, so that some goroutines
 * hold rq.workerLock.RLock(), and block at cond.Wait(). Goroutine D is
 * blocked when acquiring rq.workerLock.Lock(). Some goroutines try to process
 * jobs in queue 2, but they are blocked when acquiring rq.workerLock.RLock(),
 * since write lock has a higher priority.
 *   The fix is to not acquire rq.workerLock.RLock(), while pulling data
 * from any queue. Therefore, when a goroutine is blocked at cond.Wait(),
 * rq.workLock.RLock() is not held.
 */

package kubernetes58107

import (
	"sync"
	"testing"
	"time"
)

type RateLimitingInterface interface {
	Get()
	Put()
}

type Type struct {
	cond *sync.Cond
}

func (q *Type) Get() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.cond.Wait()
}

func (q *Type) Put() {
	q.cond.Signal()
}

type ResourceQuotaController struct {
	workerLock        sync.RWMutex
	queue             RateLimitingInterface
	missingUsageQueue RateLimitingInterface
}

func (rq *ResourceQuotaController) worker(queue RateLimitingInterface, name string) func() {
	workFunc := func() bool {
		rq.workerLock.RLock()
		defer rq.workerLock.RUnlock()
		queue.Get()
		return true
	}
	return func() {
		for {
			if quit := workFunc(); quit {
				return
			}
		}
	}
}

func (rq *ResourceQuotaController) Run() {
	go rq.worker(rq.queue, "G1")()             // G1
	go rq.worker(rq.missingUsageQueue, "G2")() // G2
}

func (rq *ResourceQuotaController) Sync() {
	for i := 0; i < 100000; i++ {
		rq.workerLock.Lock()
		time.Sleep(time.Nanosecond)
		rq.workerLock.Unlock()
	}
}

func (rq *ResourceQuotaController) HelperSignals() {
	//	time.Sleep(200 * time.Millisecond)
	for i := 0; i < 100000; i++ {
		rq.queue.Put()
		rq.missingUsageQueue.Put()
	}
}

func startResourceQuotaController() {
	resourceQuotaController := &ResourceQuotaController{
		queue:             &Type{sync.NewCond(&sync.Mutex{})},
		missingUsageQueue: &Type{sync.NewCond(&sync.Mutex{})},
	}

	go resourceQuotaController.Run()
	go resourceQuotaController.Sync() // G3
	resourceQuotaController.HelperSignals()
}

/// G1 						G2						G3
/// ...						...						Sync()
/// rq.workerLock.RLock()
/// q.cond.Wait()
/// 												rq.workerLock.Lock()
/// 						rq.workerLock.RLock()
///
func TestKubernetes58107(t *testing.T) {
	startResourceQuotaController()
}
