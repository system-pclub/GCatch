package main

import "sync"

func main() {
	mu := sync.Mutex{}
	go func() {
		mu.Lock()
		mu.Unlock()
	}()
	mu.Unlock()
}
