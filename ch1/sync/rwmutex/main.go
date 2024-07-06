package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.RWMutex
	value int
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter.Value())
}
