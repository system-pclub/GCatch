package figure1

import "sync"

type Controller struct {
	lock *sync.Mutex
	dispatch chan int
}

func main() {
	cont := Controller {
		dispatch: make(chan int),
		lock: &sync.Mutex{},
	}
	go cont.worker()
	cont.lock.Lock()
	cont.dispatch <- 100
	cont.lock.Unlock()
	cont.lock.Lock()
	cont.dispatch <- 100
	cont.lock.Unlock()
}

func (cont Controller) worker() {
	select {
	case <-cont.dispatch:
		cont.lock.Lock()
		cont.lock.Unlock()
	default:
		return
	}
	select {
	case <-cont.dispatch:
		cont.lock.Lock()
		cont.lock.Unlock()
	default:
		return
	}
}

//func main() {
//	cont := Controller {
//		dispatch: make(chan int),
//		lock: &sync.Mutex{},
//	}
//	go cont.worker()
//	for i:=0; i < 2; i++ {
//		cont.lock.Lock()
//		cont.dispatch <- 100
//		cont.lock.Unlock()
//	}
//}
//
//func (cont Controller) worker() {
//	for i := 0; i < 2; i++ {
//		select {
//		case <-cont.dispatch:
//			cont.lock.Lock()
//			cont.lock.Unlock()
//		default:
//			return
//		}
//	}
//}