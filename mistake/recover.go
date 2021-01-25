package main

import "fmt"

// 可以使用 recover ()函数捕获/拦截恐慌。只有在延迟函数中调用 recover ()时，才能实现这个功能。
func main() {
	defer func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
	}()
	panic("make error")
}

// 不能放在其他函数
func doRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出错了")
		}
	}()
}
