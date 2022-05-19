package main

func main() {
	c1 := make(chan int)
	go func() {
		close(c1)
	}()
	close(c1)
}
