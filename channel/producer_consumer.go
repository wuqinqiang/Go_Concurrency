package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go consumer(ch)
	go producer(ch)
	time.Sleep(3 * time.Second)
}

// 一个生产者
func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

// 消费者
func consumer(task <-chan int) {
	for i := 0; i < 5; i++ {
		// 5个消费者
		go func(id int) {
			for {
				item, ok := <-task
				// 如果等于false 说明通道已关闭
				if !ok {
					return
				}
				fmt.Printf("消费者:%d，消费了:%d\n", id, item)
				// 给别人一点机会不会吃亏
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}
}
