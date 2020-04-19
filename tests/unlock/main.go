package main

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
}

func (c * SafeCounter) Lock() {
	c.mux.Lock()
}

func (c * SafeCounter) Unlock() {
	c.mux.Unlock()
}

func (c * SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key] ++
	//c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func Test0() * SafeCounter {

	c := SafeCounter{v: make(map[string]int)}
	c.mux.Lock()

	return &c
}

func Test1() {
	i := 1

	for i <= 3 {
		if i == 2 {
			fmt.Println(i)
		}
		i = i + 1
	}


	if i < 10 {
		if i % 2 == 0 {
			i = i + 10
		}
	}

	fmt.Println(i)
}

func Test2() {
	i := 1

	for i <= 3 {
		if i == 2 {
			fmt.Println(i)
		}
		i = i + 1
	}

	mux := sync.Mutex{}

	j = 10

	if i < 10 {
		if j > 3 {
			mux.Lock()
		}
	}

	if i < 10 {
		if j > 3 {
			mux.Unlock()
		}
	}

	//mux1 := sync.Mutex{}


}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}