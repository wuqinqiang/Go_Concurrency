package main

import (
	"fmt"
	"time"
)

type Token struct{}

func main() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
	for i := 0; i < 4; i++ {
		go func(id int) {
			NewToken(id, chs[id], chs[(id+1)%4])
		}(i)
	}
	chs[0] <- struct{}{}
	select {}
}

func NewToken(id int, ch chan Token, nextChan chan Token) {
	for {
		res := <-ch
		fmt.Println(id + 1)
		time.Sleep(1 * time.Second)
		nextChan <- res
	}
}
