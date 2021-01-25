package main

import (
	"fmt"
	"time"
)

func main() {
	var hi string
	go func() {
		hi = "golang"
	}()
	time.Sleep(10 * time.Millisecond)
	fmt.Println(hi)
	//golang
}
