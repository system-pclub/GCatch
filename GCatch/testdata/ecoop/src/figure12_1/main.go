package figure12_1

func main() {
	c1 := make(chan int)
	go func() {
		c1 <- 1
	}()
	close(c1)
}
