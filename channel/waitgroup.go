package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	jobCount := 100
	limit := 0
	for i := 0; i < jobCount; i++ {
		limit++
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			job2(item)
		}(i)

		if limit == 10 {
			wg.Wait()
			limit = 0
		}
	}
}

func job2(index int) {
	// 耗时任务
	time.Sleep(1*time.Second)
	fmt.Printf("任务:%d已完成\n", index)
}
