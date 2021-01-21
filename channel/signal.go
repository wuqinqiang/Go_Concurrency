package main

import (
	"fmt"
	"time"
)

func main() {
	isOver := make(chan struct{})
	go func() {
		collectMsg(isOver)
	}()
	<-isOver
	calculateMsg()
}

// 采集
func collectMsg(isOver chan struct{}) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("完成采集工具")
	isOver <- struct{}{}
}

// 计算
func calculateMsg() {
	fmt.Println("开始进行数据分析")
}
