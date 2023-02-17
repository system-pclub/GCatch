package kubernetes81091

import (
	"sync"
	"testing"
)

type FakeFilterPlugin struct {
	numFilterCalled int
}

func (fp *FakeFilterPlugin) Filter() {
	fp.numFilterCalled++
}

type FilterPlugin interface {
	Filter()
}

type Framework interface {
	RunFilterPlugins()
}

type framework struct {
	filterPlugins []FilterPlugin
}

func NewFramework() Framework {
	f := &framework{}
	f.filterPlugins = append(f.filterPlugins, &FakeFilterPlugin{})
	return f
}

func (f *framework) RunFilterPlugins() {
	for _, pl := range f.filterPlugins {
		pl.Filter()
	}
}

type genericScheduler struct {
	framework Framework
}

func NewGenericScheduler(framework Framework) *genericScheduler {
	return &genericScheduler{
		framework: framework,
	}
}

func (g *genericScheduler) findNodesThatFit() {
	checkNode := func(i int) {
		g.framework.RunFilterPlugins()
	}
	ParallelizeUntil(2, 2, checkNode)
}

func (g *genericScheduler) Schedule() {
	g.findNodesThatFit()
}

type DoWorkPieceFunc func(piece int)

func ParallelizeUntil(workers, pieces int, doWorkPiece DoWorkPieceFunc) {
	var stop <-chan struct{}

	toProcess := make(chan int, pieces)
	for i := 0; i < pieces; i++ {
		toProcess <- i
	}
	close(toProcess)

	if pieces < workers {
		workers = pieces
	}

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for piece := range toProcess {
				select {
				case <-stop:
					return
				default:
					doWorkPiece(piece)
				}
			}
		}()
	}
	wg.Wait()
}

func TestKubernetes81091(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		filterFramework := NewFramework()
		scheduler := NewGenericScheduler(filterFramework)
		scheduler.Schedule()
	}()
	wg.Wait()
}
