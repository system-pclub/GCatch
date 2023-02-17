package etcd5509

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

var ErrConnClosed error

type Client struct {
	mu     sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
}

func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cancel == nil {
		return
	}
	c.cancel()
	c.cancel = nil
	c.mu.Unlock()
	c.mu.Lock() // block here
}

type remoteClient struct {
	client *Client
	mu     sync.Mutex
}

func (r *remoteClient) acquire(ctx context.Context) error {
	for {
		r.client.mu.RLock()
		closed := r.client.cancel == nil
		r.mu.Lock()
		r.mu.Unlock()
		if closed {
			return ErrConnClosed // Missing RUnlock before return
		}
		r.client.mu.RUnlock()
	}
}

type kv struct {
	rc *remoteClient
}

func (kv *kv) Get(ctx context.Context) error {
	return kv.Do(ctx)
}

func (kv *kv) Do(ctx context.Context) error {
	for {
		err := kv.do(ctx)
		if err == nil {
			return nil
		}
		return err
	}
}

func (kv *kv) do(ctx context.Context) error {
	err := kv.getRemote(ctx)
	return err
}

func (kv *kv) getRemote(ctx context.Context) error {
	return kv.rc.acquire(ctx)
}

type KV interface {
	Get(ctx context.Context) error
	Do(ctx context.Context) error
}

func NewKV(c *Client) KV {
	return &kv{rc: &remoteClient{
		client: c,
	}}
}
func TestEtcd5509(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	cli := &Client{
		ctx:    ctx,
		cancel: cancel,
	}
	kv := NewKV(cli)
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		err := kv.Get(context.TODO())
		if err != nil && err != ErrConnClosed {
			fmt.Println("Expect ErrConnClosed")
		}
	}()

	cli.Close()

	<-donec
}
