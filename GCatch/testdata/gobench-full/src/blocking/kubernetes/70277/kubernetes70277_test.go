package kubernetes70277

import (
	"testing"
	"time"
)

var ForerverTestTimeout = time.Second * 20

type WaitFunc func(done <-chan struct{}) <-chan struct{}

type ConditionFunc func() (done bool, err error)

func WaitFor(wait WaitFunc, fn ConditionFunc, done <-chan struct{}) error {
	c := wait(done)
	for {
		_, open := <-c
		ok, err := fn()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		if !open {
			break
		}
	}
	return nil
}

func poller(interval, timeout time.Duration) WaitFunc {
	return WaitFunc(func(done <-chan struct{}) <-chan struct{} {
		ch := make(chan struct{})
		go func() {
			defer close(ch)

			tick := time.NewTicker(interval)
			defer tick.Stop()

			var after <-chan time.Time
			if timeout != 0 {
				timer := time.NewTimer(timeout)
				after = timer.C
				defer timer.Stop()
			}
			for {
				select {
				case <-tick.C:
					select {
					case ch <- struct{}{}:
					default:
					}
				case <-after:
					return
				case <-done:
					return
				}
			}
		}()

		return ch
	})
}

func TestKubernetes70277(t *testing.T) {
	stopCh := make(chan struct{})
	defer close(stopCh)
	waitFunc := poller(time.Millisecond, ForerverTestTimeout)
	var doneCh <-chan struct{}

	WaitFor(func(done <-chan struct{}) <-chan struct{} {
		doneCh = done
		return waitFunc(done)
	}, func() (bool, error) {
		return true, nil
	}, stopCh)

	<-doneCh // block here
}
