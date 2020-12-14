package main

import (
	"math/rand"
	"time"
)

func gl2LeakFunc() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
		<- ch
	}()
	if rand.Intn(10) >5 {
		println("leaked a goroutine after returning")
		return
	}
	ch <- struct{}{}
}

func main() {
	gl2LeakFunc()
}