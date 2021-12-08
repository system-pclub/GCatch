package figure13_2

import "sync"

func main() {
	l1 := &sync.Mutex{}
	l2 := &sync.Mutex{}
	go func() {
		lock_both1(l1, l2)
	}()
	lock_both2(l2, l1)
}

func lock_both1(fst *sync.Mutex, sec *sync.Mutex) {
	fst.Lock()
	sec.Lock()
	sec.Unlock()
	fst.Unlock()
}

func lock_both2(fst *sync.Mutex, sec *sync.Mutex) {
	fst.Lock()
	sec.Lock()
	sec.Unlock()
	fst.Unlock()
}