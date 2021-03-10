package main

func main() {
	a := make(chan int)
	b := make(chan int)
	go func() {
		<- b
		a <- 1
	}()
	<- a
	b <- 1
}
