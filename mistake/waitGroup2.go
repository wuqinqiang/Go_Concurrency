package main

import (
	"fmt"
	"sync"
)

//是因为每个工作者都获得了原始“ WaitGroup”变量的一个副本。当工人执行 wg 时。Done ()对主 goroutine 中的“ WaitGroup”变量没有影响。
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go doJob(wg, i)
	}
	wg.Wait()
}

func doJob(wg sync.WaitGroup, item int) {
	defer wg.Done()
	fmt.Println("item is:", item)
}
