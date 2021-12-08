package figure13_1

import "sync"

func main() {
	l1 := &sync.Mutex{}
	go func() {
		l1.Lock()
	}()
	l1.Unlock()
}
