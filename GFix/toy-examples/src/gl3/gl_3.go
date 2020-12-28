package main

import (
	"math/rand"
	"time"
)

func gl3LeakFunc() {
	ch := make(chan struct{})
	go func() {
		for i := 1; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
			println("did something")
			ch <- struct{}{}
		}
	}()
	for i := 1; i < 5; i++ {
		select {
		case <-ch:
			println("received something")
		case <-time.After(5 * time.Second):
			println("possibly leaked a goroutine after returning")
			break
		}
	}
}

func main() {
	gl3LeakFunc()
}