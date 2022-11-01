package main

import (
	"sync"
)

type processorListener[T any] struct {
	lock sync.RWMutex
	cond sync.Cond

	pendingNotifications []T
}

func (p *processorListener[T]) add(notification T) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.pendingNotifications = append(p.pendingNotifications, notification)
	p.cond.Broadcast()
}

func (p *processorListener[T]) pop(stopCh <-chan struct{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for {
		for len(p.pendingNotifications) == 0 {
			select {
			case <-stopCh:
				return
			default:
			}
			p.cond.Wait()
		}
		select { // block here
		case <-stopCh:
			return
		}
	}
}

func newProcessListener[T any]() *processorListener[T] {
	ret := &processorListener[T]{
		pendingNotifications: []T{},
	}
	ret.cond.L = &ret.lock
	return ret
}
func main() {
	pl := newProcessListener[int]()
	stopCh := make(chan struct{})
	defer close(stopCh)
	pl.add(1)
	go pl.pop(stopCh)

	resultCh := make(chan struct{})
	go func() {
		pl.lock.Lock() // block here
		close(resultCh)
	}()
	<-resultCh
	pl.lock.Unlock()
}
