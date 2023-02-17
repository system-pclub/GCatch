/*
 * Project: moby
 * Issue or PR  : https://github.com/moby/moby/pull/21233
 * Buggy version: cc12d2bfaae135e63b1f962ad80e6943dd995337
 * fix commit-id: 2f4aa9658408ac72a598363c6e22eadf93dbb8a7
 * Flaky:100/100
 * Description:
 *   This test was checking that it received every progress update that was
 *  produced. But delivery of these intermediate progress updates is not
 *  guaranteed. A new update can overwrite the previous one if the previous
 *  one hasn't been sent to the channel yet.
 *    The call to t.Fatalf exited the cur rent goroutine which was consuming
 *  the channel, which caused a deadlock and eventual test timeout rather
 *  than a proper failure message.
 */
package moby21233

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

type Progress struct{}

type Output interface {
	WriteProgress(Progress) error
}

type chanOutput chan<- Progress

type TransferManager struct {
	mu sync.Mutex
}

type Transfer struct {
	mu sync.Mutex
}

type Watcher struct {
	signalChan  chan struct{}
	releaseChan chan struct{}
	running     chan struct{}
}

func ChanOutput(progressChan chan<- Progress) Output {
	return chanOutput(progressChan)
}
func (out chanOutput) WriteProgress(p Progress) error {
	out <- p
	return nil
}
func NewTransferManager() *TransferManager {
	return &TransferManager{}
}
func NewTransfer() *Transfer {
	return &Transfer{}
}
func (t *Transfer) Release(watcher *Watcher) {
	t.mu.Lock()
	t.mu.Unlock()
	close(watcher.releaseChan)
	<-watcher.running
}
func (t *Transfer) Watch(progressOutput Output) *Watcher {
	t.mu.Lock()
	defer t.mu.Unlock()
	lastProgress := Progress{}
	w := &Watcher{
		releaseChan: make(chan struct{}),
		signalChan:  make(chan struct{}),
		running:     make(chan struct{}),
	}
	go func() { // G2
		defer func() {
			close(w.running)
		}()
		done := false
		for {
			t.mu.Lock()
			t.mu.Unlock()
			if rand.Int31n(2) >= 1 {
				progressOutput.WriteProgress(lastProgress)
			}
			if done {
				return
			}
			select {
			case <-w.signalChan:
			case <-w.releaseChan:
				done = true
				select {
				default:
				}
			}
		}
	}()
	return w
}
func (tm *TransferManager) Transfer(progressOutput Output) (*Transfer, *Watcher) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	t := NewTransfer()
	return t, t.Watch(progressOutput)
}

func testTransfer() {
	tm := NewTransferManager()
	progressChan := make(chan Progress)
	progressDone := make(chan struct{})
	go func() { // G3
		for p := range progressChan { /// Chan consumer
			if rand.Int31n(2) >= 1 {
				return
			}
			fmt.Println(p)
		}
		close(progressDone)
	}()
	ids := []string{"id1", "id2", "id3"}
	xrefs := make([]*Transfer, len(ids))
	watchers := make([]*Watcher, len(ids))
	for i := range ids {
		xrefs[i], watchers[i] = tm.Transfer(ChanOutput(progressChan)) /// Chan producer
	}

	for i := range xrefs {
		xrefs[i].Release(watchers[i]) /// Block here
	}

	close(progressChan)
	<-progressDone
}

///
/// G1 						G2					G3
/// testTransfer()
/// tm.Transfer()
/// t.Watch()
/// 						WriteProgress()
/// 						ProgressChan<-
/// 											<-progressChan
/// 						...					...
/// 						return
/// 											<-progressChan
/// <-watcher.running
/// ----------------------G1, G3 leak--------------------------
///

func TestMoby21233(t *testing.T) {
	go testTransfer() // G1
}
