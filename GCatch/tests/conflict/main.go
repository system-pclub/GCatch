package conflict

import (
	"fmt"
	"sync"
	"time"
)


var j int

type SafeCounter struct {
	v map[string] int
	i int
	s string
	f float64
	mux sync.Mutex
	mux1 sync.Mutex
}

func (c * SafeCounter) Lock() {
	c.mux.Lock()
}

func (c * SafeCounter) Unlock() {
	c.mux.Unlock()
}

func (c * SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.ProtectedInc(key)
	c.mux.Unlock()
}

func (c * SafeCounter) ProtectedInc(key string) {
	c.mux1.Lock()
	c.v[key] ++
	c.mux1.Unlock()
}

func (c * SafeCounter) UnProtectedInc(key string) {
	c.v[key] ++
}

func (c *SafeCounter) Value(key string) int {
	c.mux1.Lock()
	defer c.mux1.Unlock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.mux.Lock()
	defer c.mux.Unlock()


	return c.v[key]
}





func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
