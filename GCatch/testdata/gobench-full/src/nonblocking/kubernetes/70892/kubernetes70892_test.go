package kubernetes70892

import (
	"context"
	"sync"
	"testing"
)

type HostPriorityList []int

type DoWorkPieceFunc func(piece int)

func ParallelizeUntil(ctx context.Context, workers, pieces int, doWorkPiece DoWorkPieceFunc) {
	var stop <-chan struct{}
	if ctx != nil {
		stop = ctx.Done()
	}

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

func TestKubernetes70892(t *testing.T) {
	priorityConfigs := append([]int{}, 1, 2, 3)
	results := make([]HostPriorityList, len(priorityConfigs), len(priorityConfigs))

	for i := range priorityConfigs {
		results[i] = make(HostPriorityList, 2)
	}
	processNode := func(index int) {
		for i := range priorityConfigs {
			if results[i][0] != 4 {
				results[i] = HostPriorityList{7, 8, 9}
			}
		}
	}
	ParallelizeUntil(context.Background(), 2, 2, processNode)
}
