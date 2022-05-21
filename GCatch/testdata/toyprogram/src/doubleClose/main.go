package main

type client struct {
	closed chan int
}

func newClient() *client {
	var c *client
	c = new(client)
	c.closed = make(chan int)
	return c
}

func (c *client) closeClient1() {
	select {
	case <-c.closed:
	default:
		close(c.closed)
	}
}

func (c *client) closeClient2() {
	close(c.closed)
}


func test1() {
	c := newClient()

	go func() {
		c.closeClient1()
	}()

	go func() {
		c.closeClient1()
	}()
}

func test2() {
	c := newClient()

	go func() {
		c.closeClient2()
	}()

	go func() {
		c.closeClient2()
	}()
}


func main() {
	test1()
	test2()
	c1 := make(chan int)
	go func() {
		close(c1)
	}()
	close(c1)
}
