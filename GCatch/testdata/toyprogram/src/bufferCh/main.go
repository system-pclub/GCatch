package bufferCh

func main() {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 1
}

func foo() {
	a := make(chan int, 1)
	b := make(chan int, 1)
	go func() {
		<- b
		a <- 1
	}()
	<- a
	b <- 1
}