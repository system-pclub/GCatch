/*
 * Project: kubernetes
 * Issue or PR  : https://github.com/kubernetes/kubernetes/pull/25331
 * Buggy version: 5dd087040bb13434f1ddf2f0693d0203c30f28cb
 * fix commit-id: 97f4647dc3d8cf46c2b66b89a31c758a6edfb57c
 * Flaky: 100/100
 * Description:
 *   In reflector.go, it could probably call Stop() without retrieving
 * all results from ResultChan(). See here. A potential leak is that
 * when an error has happened, it could block on resultChan, and then
 * cancelling context in Stop() wouldn't unblock it.
 */
package kubernetes25331

import (
	"context"
	"errors"
	"testing"
)

type watchChan struct {
	ctx        context.Context
	cancel     context.CancelFunc
	resultChan chan bool
	errChan    chan error
}

func (wc *watchChan) Stop() {
	wc.errChan <- errors.New("Error")
	wc.cancel()
}

func (wc *watchChan) run() {
	select {
	case err := <-wc.errChan:
		errResult := len(err.Error()) != 0
		wc.cancel() // Removed in fix
		wc.resultChan <- errResult
	case <-wc.ctx.Done():
	}
}

func NewWatchChan() *watchChan {
	ctx, cancel := context.WithCancel(context.Background())
	return &watchChan{
		ctx:        ctx,
		cancel:     cancel,
		resultChan: make(chan bool),
		errChan:    make(chan error),
	}
}

///
/// G1					G2
/// wc.run()
///						wc.Stop()
///						wc.errChan <-
///						wc.cancel()
///	<-wc.errChan
///	wc.cancel()
///	wc.resultChan <-
///	-------------G1 leak----------------
///

func TestKubernetes25331(t *testing.T) {
	wc := NewWatchChan()
	go wc.run()  // G1
	go wc.Stop() // G2
}
