package etcd9446

import (
	"sync"
	"testing"
)

type txBuffer struct {
	buckets map[string]struct{}
}

func (txb *txBuffer) reset() {
	for k, _ := range txb.buckets {
		delete(txb.buckets, k)
	}
}

type txReadBuffer struct{ txBuffer }

func (txr *txReadBuffer) Range() {
	_ = txr.buckets["1"]
}

type readTx struct {
	buf txReadBuffer
}

func (rt *readTx) reset() {
	rt.buf.reset()
}

func (rt *readTx) UnsafeRange() {
	rt.buf.Range()
}

func TestEtcd9446(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		txn := &readTx{
			buf: txReadBuffer{
				txBuffer{
					buckets: make(map[string]struct{}),
				},
			},
		}
		txn.buf.buckets["1"] = struct{}{}
		go func() {
			defer wg.Done()
			txn.reset()
		}()
		go func() {
			defer wg.Done()
			txn.UnsafeRange()
		}()
	}()
	wg.Wait()
}
