package sync

import "sync"

type Counter struct {
	sync.Mutex
	//mu sync.Mutex //mutual exclusion lock
	// zero value is unlocked
	value int
}

// Mutex: ensuring only one goroutine can increment the counter at a time

func (c *Counter) Inc() {
	//c.mu.Lock()
	//defer c.mu.Unlock()
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
