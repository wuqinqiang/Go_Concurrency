package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {

	c := sync.NewCond(&sync.Mutex{})

	var ready int

	for i := 0; i < 10; i++ {
		go func(item int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)

			c.L.Lock()
			ready++
			c.L.Unlock()
			fmt.Printf("第%d位运动员准备就绪\n", item)
			c.Broadcast()
		}(i)
	}

	c.L.Lock()

	for ready != 10 {
		c.Wait()
		fmt.Println("裁判被唤醒一次")
	}
	c.L.Unlock()
	log.Println("所有运动员都准备就绪。比赛开始，3，2，1, ......")
}
