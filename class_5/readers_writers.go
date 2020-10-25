package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				res := count.Count() // 计数器读操作
				fmt.Println("res的值是:", res)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		count.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}

type Counter struct {
	mu    sync.RWMutex
	count int64
}

func (c *Counter) Count() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}
