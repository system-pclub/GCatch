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

func (c *client) closeClient() {
	select {
	case <-c.closed:
	default:
		close(c.closed)
	}
}


func test1() {
	c := newClient()

	go func() {
		c.closeClient()
	}()

	go func() {
		c.closeClient()
	}()
}


func main() {
	test1()
	c1 := make(chan int)
	go func() {
		close(c1)
	}()
	close(c1)
}
