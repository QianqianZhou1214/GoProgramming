package sync

import "sync"

type Counter struct {
	mu sync.Mutex //mutual exclusion lock
	// zero value is unlocked
	value int
}

// do not embedding types, otherwise the methods will become part of public interface

// Mutex: ensuring only one goroutine can increment the counter at a time

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
