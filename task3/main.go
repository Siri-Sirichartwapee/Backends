package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc เพิ่มค่า count ทีละ 1
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Value คืนค่าปัจจุบัน
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	var counter SafeCounter
	var wg sync.WaitGroup

	// จำลอง 1000 goroutines เรียก Inc พร้อมกัน
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}

	wg.Wait()

	fmt.Println("Final count =", counter.Value())
}