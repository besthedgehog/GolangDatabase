package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	likes int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.likes++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.likes
}

func main() {
	counter := &SafeCounter{}

	var wg sync.WaitGroup

	for _ = range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Println("Всего лайков")
	fmt.Println(counter.Value())

	// Обходим защиту
	counter.Inc()
	fmt.Println(counter.Value())
}
