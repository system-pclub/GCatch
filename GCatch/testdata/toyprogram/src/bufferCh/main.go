package bufferCh

func main() {
	ch := make(chan int, 1)
	ch <- 1
	ch <- 1
}
