package figure12_3

func main() {
	c1 := make(chan int)
	go func() {
		<-c1
	}()
}
