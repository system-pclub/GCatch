/*
 * Project: kubernetes
 * Issue or PR  : https://github.com/kubernetes/kubernetes/pull/5316
 * Buggy version: c868b0bbf09128960bc7c4ada1a77347a464d876
 * fix commit-id: cc3a433a7abc89d2f766d4c87eaae9448e3dc091
 * Flaky: 100/100
 * Description:
 *   If the main goroutine selects a case that doesn’t consumes
 * the channels, the anonymous goroutine will be blocked on sending
 * to channel.
 */

package kubernetes5316

import (
	"errors"
	"math/rand"
	"testing"
	"time"
)

func finishRequest(timeout time.Duration, fn func() error) {
	ch := make(chan bool)     // FIX: ch := make(chan bool, 1)
	errCh := make(chan error) // FIX: errCh := make(chan error, 1)
	go func() {               // G2
		if err := fn(); err != nil {
			errCh <- err
		} else {
			ch <- true
		}
	}()

	select {
	case <-ch:
	case <-errCh:
	case <-time.After(timeout):
	}
}

///
/// G1 						G2
/// finishRequest()
/// 						fn()
/// time.After()
/// 						errCh<-/ch<-
/// --------------G2 leak----------------
///

func TestKubernetes5316(t *testing.T) {
	fn := func() error {
		time.Sleep(2 * time.Millisecond)
		if rand.Intn(10) > 5 {
			return errors.New("Error")
		}
		return nil
	}
	go finishRequest(time.Millisecond, fn) // G1
}
