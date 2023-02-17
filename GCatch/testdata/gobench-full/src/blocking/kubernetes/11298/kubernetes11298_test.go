package kubernetes11298

import (
	"sync"
	"testing"
	"time"
)

type Signal <-chan struct{}

func After(f func()) Signal {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		if f != nil {
			f()
		}
	}()
	return Signal(ch)
}

func Until(f func(), period time.Duration, stopCh <-chan struct{}) {
	if f == nil {
		return
	}
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
		case <-time.After(period):
		}
	}

}

type notifier struct {
	lock sync.Mutex
	cond *sync.Cond
}

func (n *notifier) serviceLoop(abort <-chan struct{}) {
	n.lock.Lock()
	defer n.lock.Unlock()
	for {
		select {
		case <-abort:
			return
		default:
			ch := After(n.cond.Wait)
			select {
			case <-abort:
				n.cond.Signal()
				<-ch
				return
			case <-ch:
			}
		}
	}
}

func Notify(abort <-chan struct{}) {
	n := &notifier{}
	n.cond = sync.NewCond(&n.lock)
	finished := After(func() {
		Until(func() {
			for {
				select {
				case <-abort:
					return
				default:
					func() {
						n.lock.Lock()
						defer n.lock.Unlock()
						n.cond.Signal()
					}()
				}
			}
		}, 0, abort)
	})
	Until(func() { n.serviceLoop(finished) }, 0, abort)
}
func TestKubernetes11298(t *testing.T) {
	done := make(chan struct{})
	notifyDone := After(func() { Notify(done) })
	go func() {
		defer close(done)
		time.Sleep(300 * time.Nanosecond)
	}()
	<-notifyDone
}
