package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)
	child := context.WithValue(ctx, "name", "wuqq")
	go func() {
		for {
			select {
			case <-child.Done():
				fmt.Println("it's over")
				return
			default:
				res := child.Value("name")
				fmt.Println("name:", res)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()

	time.Sleep(5 * time.Second)
}
