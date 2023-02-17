package istio8144

import (
	"sync"
	"testing"
)

type EvictionCallback func()

type callbackRecorder struct {
	callbacks int
}

func (c *callbackRecorder) callback() {
	c.callbacks++
}

type ttlCache struct {
	entries  sync.Map
	callback func()
}

func (c *ttlCache) evicter() {
	c.evictExpired()
}

func (c *ttlCache) evictExpired() {
	c.entries.Range(func(key interface{}, value interface{}) bool {
		c.callback()
		return true
	})
}

func (c *ttlCache) SetWithExpiration(key interface{}, value interface{}) {
	c.entries.Store(key, value)
}

func NewTTLWithCallback(callback EvictionCallback) *ttlCache {
	c := &ttlCache{
		callback: callback,
	}
	go c.evicter()
	return c
}

func TestIstio8144(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		c := &callbackRecorder{callbacks: 0}
		ttl := NewTTLWithCallback(c.callback)
		ttl.SetWithExpiration(1, 1)
		if c.callbacks != 1 {
		}
	}()
	wg.Wait()
}
