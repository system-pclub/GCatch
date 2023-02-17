/*
 * Project: cockroach
 * Issue or PR  : https://github.com/cockroachdb/cockroach/pull/13197
 * Buggy version: fff27aedabafe20cef57f75905fe340cab48c2a4
 * fix commit-id: 9bf770cd8f6eaff5441b80d3aec1a5614e8747e1
 * Flaky: 100/100
 * Description: One goroutine executing (*Tx).awaitDone() blocks and
 * waiting for a signal context.Done().
 */
package cockroach13197

import (
	"context"
	"testing"
)

type DB struct{}

func (db *DB) begin(ctx context.Context) *Tx {
	ctx, cancel := context.WithCancel(ctx)
	tx := &Tx{
		cancel: cancel,
		ctx:    ctx,
	}
	go tx.awaitDone() // G2
	return tx
}

type Tx struct {
	cancel context.CancelFunc
	ctx    context.Context
}

func (tx *Tx) awaitDone() {
	<-tx.ctx.Done()
}

func (tx *Tx) Rollback() {
	tx.rollback()
}

func (tx *Tx) rollback() {
	tx.close()
}

func (tx *Tx) close() {
	tx.cancel()
}

/// G1 				G2
/// begin()
/// 				awaitDone()
/// 				<-tx.ctx.Done()
/// return
/// -----------G2 leak-------------
func TestCockroach13197(t *testing.T) {
	db := &DB{}
	db.begin(context.Background()) // G1
}
