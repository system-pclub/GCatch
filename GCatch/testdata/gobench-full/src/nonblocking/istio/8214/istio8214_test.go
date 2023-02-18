package istio8214

import (
	"sync"
	"sync/atomic"
	"testing"
)

type internal_Cache interface {
	Set()
	Stats() Stats
}

type ExpiringCache interface {
	internal_Cache
	SetWithExpiration()
}

type Cache struct {
	cache ExpiringCache
}

func (cc *Cache) Set() {
	cc.cache.SetWithExpiration()
	cc.recordStats()
}

func (cc *Cache) recordStats() {
	cc.cache.Stats()
}

type Stats struct {
	Writes uint64
}

type lruCache struct {
	stats Stats
}

func (c *lruCache) Stats() Stats {
	return c.stats
}

func (c *lruCache) Set() {
	c.SetWithExpiration()
}

func (c *lruCache) SetWithExpiration() {
	atomic.AddUint64(&c.stats.Writes, 1)
}

type grpcServer struct {
	cache *Cache
}

func (s *grpcServer) check() {
	if s.cache != nil {
		s.cache.Set()
	}
}

func (s *grpcServer) Check() {
	s.check()
}

func TestIstio8214(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		s := &grpcServer{
			cache: &Cache{
				cache: &lruCache{},
			},
		}
		go func() {
			defer wg.Done()
			s.Check()
		}()
		go func() {
			defer wg.Done()
			s.Check()
		}()
	}()
	wg.Wait()
}
