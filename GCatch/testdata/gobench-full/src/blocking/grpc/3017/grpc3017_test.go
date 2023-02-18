package grpc3017

import (
	"sync"
	"testing"
	"time"
)

type Address int
type SubConn int

type subConnCacheEntry struct {
	sc            SubConn
	cancel        func()
	abortDeleting bool
}

type lbCacheClientConn struct {
	mu            sync.Mutex
	timeout       time.Duration
	subConnCache  map[Address]*subConnCacheEntry
	subConnToAddr map[SubConn]Address
}

func (ccc *lbCacheClientConn) NewSubConn(addrs []Address) SubConn {
	if len(addrs) != 1 {
		return SubConn(1)
	}
	addrWithoutMD := addrs[0]
	ccc.mu.Lock()
	defer ccc.mu.Unlock()
	if entry, ok := ccc.subConnCache[addrWithoutMD]; ok {
		entry.cancel()
		delete(ccc.subConnCache, addrWithoutMD)
		return entry.sc
	}
	scNew := SubConn(1)
	ccc.subConnToAddr[scNew] = addrWithoutMD
	return scNew
}

func (ccc *lbCacheClientConn) RemoveSubConn(sc SubConn) {
	ccc.mu.Lock()
	defer ccc.mu.Unlock()
	addr, ok := ccc.subConnToAddr[sc]
	if !ok {
		return
	}

	if entry, ok := ccc.subConnCache[addr]; ok {
		if entry.sc != sc {
			delete(ccc.subConnToAddr, sc)
		}
		return
	}

	entry := &subConnCacheEntry{
		sc: sc,
	}
	ccc.subConnCache[addr] = entry

	timer := time.AfterFunc(ccc.timeout, func() {
		ccc.mu.Lock()
		if entry.abortDeleting {
			return // Missing unlock
		}
		delete(ccc.subConnToAddr, sc)
		delete(ccc.subConnCache, addr)
		ccc.mu.Unlock()
	})

	entry.cancel = func() {
		if !timer.Stop() {
			entry.abortDeleting = true
		}
	}
}
func TestGrpc3017(t *testing.T) {
	done := make(chan struct{})

	ccc := &lbCacheClientConn{
		timeout:       time.Nanosecond,
		subConnCache:  make(map[Address]*subConnCacheEntry),
		subConnToAddr: make(map[SubConn]Address),
	}

	sc := ccc.NewSubConn([]Address{Address(1)})
	go func() {
		for i := 0; i < 1000; i++ {
			ccc.RemoveSubConn(sc)
			sc = ccc.NewSubConn([]Address{Address(1)})
		}
		close(done)
	}()
	<-done
}
