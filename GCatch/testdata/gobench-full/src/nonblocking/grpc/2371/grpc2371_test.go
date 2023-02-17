package grpc2371

import (
	"sync"
	"testing"
	"time"
)

type ccBalancerWrapper struct {
	cc               *ClientConn
	resolverUpdateCh chan struct{}
}

func (ccb *ccBalancerWrapper) handleResolvedAddrs() {
	select {
	case <-ccb.resolverUpdateCh:
	default:
	}
	ccb.resolverUpdateCh <- struct{}{}
}

func newCCBalancerWrapper(cc *ClientConn) *ccBalancerWrapper {
	ccb := &ccBalancerWrapper{
		cc:               cc,
		resolverUpdateCh: make(chan struct{}, 1),
	}
	return ccb
}

type ccResolverWrapper struct {
	cc *ClientConn
}

func (ccr *ccResolverWrapper) start() {
	go ccr.watcher()
}

func (ccr *ccResolverWrapper) watcher() {
	ccr.cc.handleServiceConfig()
}

func newCCResolverWrapper(cc *ClientConn) *ccResolverWrapper {
	ccr := &ccResolverWrapper{
		cc: cc,
	}
	return ccr
}

type ClientConn struct {
	mu              sync.RWMutex
	balancerWrapper *ccBalancerWrapper
	resolverWrapper *ccResolverWrapper
}

func (cc *ClientConn) handleServiceConfig() {
	cc.mu.Lock()
	cc.balancerWrapper.handleResolvedAddrs()
	cc.mu.Unlock()
}

func (cc *ClientConn) Close() {
	cc.mu.Lock()
	cc.resolverWrapper = nil
	cc.balancerWrapper = nil
	cc.mu.Unlock()
}

func Dial() *ClientConn {
	return DialContext()
}

func DialContext() *ClientConn {
	cc := &ClientConn{}

	cc.resolverWrapper = newCCResolverWrapper(cc)
	cc.balancerWrapper = newCCBalancerWrapper(cc)

	cc.resolverWrapper.start()

	return cc
}

func TestGrpc2371(t *testing.T) {

	for i := 0; i < 10; i++ {
		cc := Dial()

		go cc.Close()
	}

	time.Sleep(100 * time.Millisecond)
}
