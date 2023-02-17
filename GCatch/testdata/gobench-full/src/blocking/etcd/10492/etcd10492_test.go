package etcd10492

import (
	"context"
	"sync"
	"testing"
	"time"
)

type Checkpointer func(ctx context.Context)

type lessor struct {
	mu                 sync.RWMutex
	cp                 Checkpointer
	checkpointInterval time.Duration
}

func (le *lessor) Checkpoint() {
	le.mu.Lock() // block here
	defer le.mu.Unlock()
}

func (le *lessor) SetCheckpointer(cp Checkpointer) {
	le.mu.Lock()
	defer le.mu.Unlock()

	le.cp = cp
}

func (le *lessor) Renew() {
	le.mu.Lock()
	unlock := func() { le.mu.Unlock() }
	defer func() { unlock() }()

	if le.cp != nil {
		le.cp(context.Background())
	}
}
func TestEtcd10492(t *testing.T) {
	le := &lessor{
		checkpointInterval: 0,
	}
	fakerCheckerpointer := func(ctx context.Context) {
		le.Checkpoint()
	}
	le.SetCheckpointer(fakerCheckerpointer)
	le.mu.Lock()
	le.mu.Unlock()
	le.Renew()
}
