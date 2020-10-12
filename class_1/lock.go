package main

import (
	"fmt"
	"sync"
)

//

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var count = 0

	// wg.Add(10)
	for i := 0; i < 10; i++ {
		// 大部分时候，我们是不知道总数是多少的
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("count的值是:", count)
}
