package main

import (
	"fmt"
	"sync"
)

func main() {
	// 发送 count++ 信号channel
	ch := make(chan struct{})

	// 计数完毕通知的chan
	closeCh := make(chan struct{})

	var count = 0

	// 等待
	var wg sync.WaitGroup

	go func() {
		// 接收者从ch通道中接收信号值，自增count
		for _ = range ch {
			count++
		}
		closeCh <- struct{}{}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				// 发送着发送信号到通道告知通道接收者自增count
				ch <- struct{}{}
			}
		}()
	}
	// 等待所有g运行结束
	wg.Wait()
	// 关闭发送通道
	close(ch)
	// count计数完毕
	<-closeCh
	close(closeCh)
	fmt.Println("count的值是:", count)
}
