package main

import (
	"math/rand"
	"time"
)

func gl1LeakFunc() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		println("not leaked")
	case <-time.After(5 * time.Second):
		println("leaked a goroutine after returning")
	}
}

func main() {
	gl1LeakFunc()
}