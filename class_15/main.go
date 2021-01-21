package main

import "sync"

var mu sync.Mutex
var s string

func foo() {
	s = "hello, world"
	mu.Unlock()
}

func main() {
	mu.Lock()
	go foo()
	mu.Lock()
	print(s)
}
