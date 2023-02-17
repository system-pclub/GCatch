package serving2137

import (
	"runtime"
	"sync"
	"testing"
)

type token struct{}

type request struct {
	lock     *sync.Mutex
	accepted chan bool
}

type Breaker struct {
	pendingRequests chan token
	activeRequests  chan token
}

func (b *Breaker) Maybe(thunk func()) bool {
	var t token
	select {
	default:
		// Pending request queue is full.  Report failure.
		return false
	case b.pendingRequests <- t:
		// Pending request has capacity.
		// Wait for capacity in the active queue.
		b.activeRequests <- t
		// Defer releasing capacity in the active and pending request queue.
		defer func() { <-b.activeRequests; <-b.pendingRequests }()
		// Do the thing.
		thunk()
		// Report success
		return true
	}
}

func (b *Breaker) concurrentRequest() request {
	runtime.Gosched()

	r := request{lock: &sync.Mutex{}, accepted: make(chan bool, 1)}
	r.lock.Lock()
	var start sync.WaitGroup
	start.Add(1)
	go func() { // G2, G3
		start.Done()
		ok := b.Maybe(func() {
			r.lock.Lock() // Will block on locked mutex.
			r.lock.Unlock()
		})
		r.accepted <- ok
	}()
	start.Wait() // Ensure that the go func has had a chance to execute.
	return r
}

// Perform n requests against the breaker, returning mutexes for each
// request which succeeded, and a slice of bools for all requests.
func (b *Breaker) concurrentRequests(n int) []request {
	requests := make([]request, n)
	for i := range requests {
		requests[i] = b.concurrentRequest()
	}
	return requests
}

func NewBreaker(queueDepth, maxConcurrency int32) *Breaker {
	return &Breaker{
		pendingRequests: make(chan token, queueDepth+maxConcurrency),
		activeRequests:  make(chan token, maxConcurrency),
	}
}

func unlock(req request) {
	req.lock.Unlock()
	// Verify that function has completed
	ok := <-req.accepted
	// Requeue for next usage
	req.accepted <- ok
}

func unlockAll(requests []request) {
	for _, lc := range requests {
		unlock(lc)
	}
}

//
// G1                           G2                      G3
// b.concurrentRequests(2)
// b.concurrentRequest()
// r.lock.Lock()
//                                                      start.Done()
// start.Wait()
// b.concurrentRequest()
// r.lock.Lock()
//                              start.Done()
// start.Wait()
// unlockAll(locks)
// unlock(lc)
// req.lock.Unlock()
// ok := <-req.accepted
//                              b.Maybe()
//                              b.activeRequests <- t
//                              thunk()
//                              r.lock.Lock()
//                                                      b.Maybe()
//                                                      b.activeRequests <- t
// ----------------------------G1,G2,G3 deadlock-----------------------------
//
func TestServing2137(t *testing.T) {
	b := NewBreaker(1, 1)

	locks := b.concurrentRequests(2) // G1
	unlockAll(locks)
}
