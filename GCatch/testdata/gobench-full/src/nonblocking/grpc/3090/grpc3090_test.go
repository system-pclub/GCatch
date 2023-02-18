package grpc3090

import (
	"sync"
	"testing"
	"time"
)

type resolver_ClientConn interface {
	UpdateState()
}

type resolver_Resolver struct {
	CC resolver_ClientConn
}

func (r *resolver_Resolver) Build(cc resolver_ClientConn) Resolver {
	r.CC = cc
	r.UpdateState()
	return r
}

func (r *resolver_Resolver) ResolveNow() {
}

func (r *resolver_Resolver) UpdateState() {
	r.CC.UpdateState()
}

type Resolver interface {
	ResolveNow()
}

type ccResolverWrapper struct {
	cc       *ClientConn
	resolver Resolver
	mu       sync.Mutex
}

func (ccr *ccResolverWrapper) resolveNow() {
	ccr.mu.Lock()
	ccr.resolver.ResolveNow()
	ccr.mu.Unlock()
}

func (ccr *ccResolverWrapper) poll() {
	ccr.mu.Lock()
	defer ccr.mu.Unlock()
	go func() {
		ccr.resolveNow()
	}()
}

func (ccr *ccResolverWrapper) UpdateState() {
	ccr.poll()
}

func newCCResolverWrapper(cc *ClientConn) {
	rb := cc.dopts.resolverBuilder
	ccr := &ccResolverWrapper{}
	ccr.resolver = rb.Build(ccr)
}

type Builder interface {
	Build(cc resolver_ClientConn) Resolver
}

type dialOptions struct {
	resolverBuilder Builder
}

type ClientConn struct {
	dopts dialOptions
}

func DialContext() {
	cc := &ClientConn{
		dopts: dialOptions{},
	}
	if cc.dopts.resolverBuilder == nil {
		cc.dopts.resolverBuilder = &resolver_Resolver{}
	}
	newCCResolverWrapper(cc)
}
func Dial() {
	DialContext()
}

func TestGrpc3090(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Dial()
		time.Sleep(5 * time.Millisecond)
	}()
	wg.Wait()
}
