package exercise5

import (
	"sync"
)

func NewCounter() *Counter {
	return &Counter{}
}

type Counter struct {
	sync.Mutex
	value int
}

func (c *Counter) Inc() {
	// Comment those 2 lines and check tests with and without -race flag
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Dec() {
	c.Lock()
	defer c.Unlock()
	c.value--
}

func (c *Counter) Reset() {
	c.Lock()
	defer c.Unlock()
	c.value = 0
}

func (c *Counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.value
}
