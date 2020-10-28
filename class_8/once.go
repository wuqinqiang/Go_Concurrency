package main

import (
	"fmt"
	"sync"
)

func Other(o *sync.Once) {
	o.Do(func() {
		fmt.Println("这样当然也不行")
	})
}

func main() {
	var once sync.Once

	fun1 := func() {
		fmt.Println("第一次打印")
	}
	once.Do(fun1)

	fun2 := func() {
		fmt.Println("第二次打印")
	}
	Other(&once)
	once.Do(fun2)
}
