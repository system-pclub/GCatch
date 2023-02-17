/*
 * Project: moby
 * Issue or PR  : https://github.com/moby/moby/pull/4395
 * Buggy version: 6d6ec5e0051ad081be3d71e20b39a25c711b4bc3
 * fix commit-id: d3a6ee1e55a53ee54b91ffb6c53ba674768cf9de
 * Flaky: 100/100
 * Description:
 *   The anonyous goroutine could be waiting on sending to
 * the channel which might never be drained.
 */

package moby4395

import (
	"errors"
	"testing"
)

func Go(f func() error) chan error {
	ch := make(chan error)
	go func() {
		ch <- f() // G2
	}()
	return ch
}

///
/// G1				G2
/// Go()
/// return ch
/// 				ch <- f()
/// ----------G2 leak-------------
///

func TestMoby4395(t *testing.T) {
	Go(func() error { // G1
		return errors.New("")
	})
}
