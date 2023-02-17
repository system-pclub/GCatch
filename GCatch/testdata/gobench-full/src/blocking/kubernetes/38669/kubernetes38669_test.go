package kubernetes38669

import (
	"sync"
	"testing"
)

type Event int
type watchCacheEvent int

type cacheWatcher struct {
	sync.Mutex
	input   chan watchCacheEvent
	result  chan Event
	stopped bool
}

func (c *cacheWatcher) process(initEvents []watchCacheEvent) {
	for _, event := range initEvents {
		c.sendWatchCacheEvent(&event)
	}
	defer close(c.result)
	defer c.Stop()
	for {
		_, ok := <-c.input
		if !ok {
			return
		}
	}
}

func (c *cacheWatcher) sendWatchCacheEvent(event *watchCacheEvent) {
	c.result <- Event(*event)
}

func (c *cacheWatcher) Stop() {
	c.stop()
}

func (c *cacheWatcher) stop() {
	c.Lock()
	defer c.Unlock()
	if !c.stopped {
		c.stopped = true
		close(c.input)
	}
}

func newCacheWatcher(chanSize int, initEvents []watchCacheEvent) *cacheWatcher {
	watcher := &cacheWatcher{
		input:   make(chan watchCacheEvent, chanSize),
		result:  make(chan Event, chanSize),
		stopped: false,
	}
	go watcher.process(initEvents)
	return watcher
}

func TestKubernetes38669(t *testing.T) {
	initEvents := []watchCacheEvent{1, 2}
	w := newCacheWatcher(0, initEvents)
	w.Stop()
}
