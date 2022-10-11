package etcd7443

import (
	"context"
	"sync"
	"testing"
)

type addrConn struct {
	mu    sync.Mutex
	cc    *ClientConn
	addr  Address
	dopts dialOptions
	down  func()
}

func (ac *addrConn) tearDown() {
	ac.mu.Lock()
	defer ac.mu.Unlock()
	if ac.down != nil {
		ac.down()
		ac.down = nil
	}
}

func (ac *addrConn) resetTransport() {
	ac.mu.Lock()
	if ac.cc.dopts.balancer != nil {
		ac.down = ac.cc.dopts.balancer.Up(ac.addr)
	}
	ac.mu.Unlock()
}

type ClientConn struct {
	dopts dialOptions
	mu    sync.RWMutex
	conns map[Address]*addrConn
}

func (cc *ClientConn) lbWatcher() {
	for addrs := range cc.dopts.balancer.Notify() {
		var (
			add []Address
			del []*addrConn
		)
		cc.mu.Lock()
		for _, a := range addrs {
			if _, ok := cc.conns[a]; !ok {
				add = append(add, a)
			}
		}

		for k, c := range cc.conns {
			var keep bool
			for _, a := range addrs {
				if k == a {
					keep = true
					break
				}
			}
			if !keep {
				del = append(del, c)
				delete(cc.conns, c.addr)
			}
		}
		cc.mu.Unlock()
		for _, a := range add {
			cc.resetAddrConn(a)
		}
		for _, c := range del {
			c.tearDown()
		}
	}
}

func (cc *ClientConn) resetAddrConn(addr Address) {
	ac := &addrConn{
		cc:    cc,
		addr:  addr,
		dopts: cc.dopts,
	}
	cc.mu.Lock()
	if cc.conns == nil {
		cc.mu.Unlock()
		return
	}
	cc.conns[ac.addr] = ac
	cc.mu.Unlock()
	go func() {
		ac.resetTransport()
	}()
}

func (cc *ClientConn) Close() {
	cc.mu.Lock()
	conns := cc.conns
	cc.conns = nil
	cc.mu.Unlock()
	if cc.dopts.balancer != nil {
		cc.dopts.balancer.Close()
	}
	for _, ac := range conns {
		ac.tearDown()
	}
}

type dialOptions struct {
	balancer Balancer
}
type DialOption func(*dialOptions)

func Dial(opts ...DialOption) *ClientConn {
	return DialContext(context.Background(), opts...)
}

func DialContext(ctx context.Context, opts ...DialOption) *ClientConn {
	cc := &ClientConn{
		conns: make(map[Address]*addrConn),
	}
	for _, opt := range opts {
		opt(&cc.dopts)
	}
	go cc.lbWatcher()
	return cc
}

type Balancer interface {
	Up(addr Address) (down func())
	Notify() <-chan []Address
	Close()
}

type Address int

type simpleBalancer struct {
	addrs    []Address
	notifyCh chan []Address
	mu       sync.RWMutex
	closed   bool
	pinAddr  Address
}

func (b *simpleBalancer) Up(addr Address) func() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return func() {}
	}

	if b.pinAddr == 0 {
		b.pinAddr = addr
		b.notifyCh <- []Address{addr}
	}

	return func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		b.mu.Lock()
		defer b.mu.Unlock()
		if b.pinAddr == addr {
			b.pinAddr = 0
			b.notifyCh <- b.addrs
		}
	}
}

func (b *simpleBalancer) Notify() <-chan []Address {
	return b.notifyCh
}

func (b *simpleBalancer) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.closed {
		return
	}
	b.closed = true
	close(b.notifyCh)
	b.pinAddr = 0
}

func newSimpleBalancer() *simpleBalancer {
	notifyCh := make(chan []Address, 1)
	addrs := make([]Address, 3)
	for i := 0; i < 3; i++ {
		addrs[i] = Address(i)
	}
	notifyCh <- addrs
	return &simpleBalancer{
		addrs:    addrs,
		notifyCh: notifyCh,
	}
}

func WithBalancer(b Balancer) DialOption {
	return func(o *dialOptions) {
		o.balancer = b
	}
}
func TestEtcd7443(t *testing.T) {
	sb := newSimpleBalancer()
	conn := Dial(WithBalancer(sb))

	closec := make(chan struct{})
	go func() {
		defer close(closec)
		sb.Close()
	}()
	go conn.Close()
	<-closec
}
