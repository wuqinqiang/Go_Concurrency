package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan struct{}) // 发送和接收count++"信号"的通道

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				ch <- struct{}{}
			}
		}()
	}

	go func() {
		wg.Wait() // 等待上面所有的 goroutine 运行完成
		close(ch) // 关闭ch通道
	}()

	count := 0
	for range ch { // 如果ch通道读取完了(ch是关闭状态), 则for循环结束
		count++
	}
	fmt.Println("count的值是:", count)
}
