package main

import "sync"

func main() {
	var mu sync.Mutex
	mu.Lock()

}

func doSomething(l sync.Mutex) {
	l.Unlock()
}
