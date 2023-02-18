/*
 * Project: etcd
 * Issue or PR  : https://github.com/etcd-io/etcd/commit/7618fdd1d642e47cac70c03f637b0fd798a53a6e
 * Buggy version: 377f19b0031f9c0aafe2aec28b6f9019311f52f9
 * fix commit-id: 7618fdd1d642e47cac70c03f637b0fd798a53a6e
 * Flaky: 9/100
 */
package etcd6873

import (
	"sync"
	"testing"
)

type watchBroadcast struct{}

type watchBroadcasts struct {
	mu      sync.Mutex
	updatec chan *watchBroadcast
	donec   chan struct{}
}

func newWatchBroadcasts() *watchBroadcasts {
	wbs := &watchBroadcasts{
		updatec: make(chan *watchBroadcast, 1),
		donec:   make(chan struct{}),
	}
	go func() { // G2
		defer close(wbs.donec)
		for wb := range wbs.updatec {
			wbs.coalesce(wb)
		}
	}()
	return wbs
}

func (wbs *watchBroadcasts) coalesce(wb *watchBroadcast) {
	wbs.mu.Lock()
	wbs.mu.Unlock()
}

func (wbs *watchBroadcasts) stop() {
	wbs.mu.Lock()
	defer wbs.mu.Unlock()
	close(wbs.updatec)
	<-wbs.donec
}

func (wbs *watchBroadcasts) update(wb *watchBroadcast) {
	select {
	case wbs.updatec <- wb:
	default:
	}
}

///
/// G1						G2					G3
/// newWatchBroadcasts()
///	wbs.update()
/// wbs.updatec <-
/// return
///							<-wbs.updatec
///							wbs.coalesce()
///												wbs.stop()
///												wbs.mu.Lock()
///												close(wbs.updatec)
///												<-wbs.donec
///							wbs.mu.Lock()
///---------------------G2,G3 deadlock-------------------------
///
func TestEtcd(t *testing.T) {
	wbs := newWatchBroadcasts() // G1
	wbs.update(&watchBroadcast{})
	go wbs.stop() // G3
}
