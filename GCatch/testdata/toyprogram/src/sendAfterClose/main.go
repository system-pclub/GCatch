package main

func main() {
	c1 := make(chan int)
	go func() {
		c1 <- 1
	}()
	close(c1)
}
