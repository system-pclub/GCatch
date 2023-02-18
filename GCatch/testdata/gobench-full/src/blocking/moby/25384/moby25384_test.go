/*
 * Project: moby
 * Issue or PR  : https://github.com/moby/moby/pull/25384
 * Buggy version: 58befe3081726ef74ea09198cd9488fb42c51f51
 * fix commit-id: 42360d164b9f25fb4b150ef066fcf57fa39559a7
 * Flaky: 100/100
 * Description:
 *   When n=1 (len(pm.plugins)), the location of group.Wait() doesn’t matter.
 * When n is larger than 1, group.Wait() is invoked in each iteration. Whenever
 * group.Wait() is invoked, it waits for group.Done() to be executed n times.
 * However, group.Done() is only executed once in one iteration.
 */
package moby25384

import (
	"sync"
	"testing"
)

type plugin struct{}

type Manager struct {
	plugins []*plugin
}

func (pm *Manager) init() {
	var group sync.WaitGroup
	group.Add(len(pm.plugins))
	for _, p := range pm.plugins {
		go func(p *plugin) {
			defer group.Done()
		}(p)
		group.Wait() // Block here
	}
}
func TestMoby25384(t *testing.T) {
	p1 := &plugin{}
	p2 := &plugin{}
	pm := &Manager{
		plugins: []*plugin{p1, p2},
	}
	go pm.init()
}
