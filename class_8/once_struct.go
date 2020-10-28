package main

import (
	"fmt"
	"sync"
	"time"
)

var threeOnce struct {
	v int
	sync.Once
}

func main() {
	for i := 3; i < 100; i++ {
		go func(item int) {
			threeOnce.Do(func() {
				fmt.Println("先进来的值:", item)
				threeOnce.v += item
			})
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("v:", threeOnce.v)
}
