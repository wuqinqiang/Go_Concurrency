package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("业务处理完毕")
	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	doClearUp()
	fmt.Println("完美谢幕")
}

func doClearUp() {
	time.Sleep(3 * time.Second)
	fmt.Println("清除任务完成")
}
