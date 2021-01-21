package main

import (
	"fmt"
	"time"
)

func main() {
	limit := make(chan struct{}, 10)
	jobCount := 100
	for i := 0; i < jobCount; i++ {
		go func(index int) {
			limit <- struct{}{}
			job(index)
			<-limit
		}(i)
	}
	time.Sleep(20 * time.Second)
}

func job(index int) {
	// 耗时任务
	time.Sleep(1 * time.Second)
	fmt.Printf("任务:%d已完成\n", index)
}
