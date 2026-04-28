package main

import (
	"fmt"
	"sync"
)

// SafeCounter is a thread-safe counter.
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc increments the counter by 1.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Value returns the current value of the counter.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	fmt.Println("--- Task 2: Safe Counter ---")
	counter := SafeCounter{}
	var wg sync.WaitGroup

	numRoutines := 1000
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d (Expected: %d)\n", counter.Value(), numRoutines)
}
