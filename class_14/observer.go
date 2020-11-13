package main

import (
	"fmt"
)

func main() {
	cha1 := make(chan struct{}, 1)

	over := make(chan struct{}, 1)
	listens := []chan struct{}{make(chan struct{}, 1), make(chan struct{}, 1), make(chan struct{}, 1)}
	fanOut(cha1, listens, true)

	cha1 <- struct{}{}
	cha1 <- struct{}{}
	cha1 <- struct{}{}

	close(cha1)
	go func() {
		for i := 0; i < len(listens); i++ {
			item := i
			go func(item int) {
				for {
					_, ok := <-listens[item]
					if !ok {
						over <- struct{}{}
						close(over)
						return
					}
					fmt.Printf("第%d监听者收到信息\n", item)
				}

			}(item)
		}
	}()

	<-over
}

func fanOut(ch1 <-chan struct{}, listens []chan struct{}, isSync bool) {
	go func() {
		defer func() {
			for i := 0; i < len(listens); i++ {
				close(listens[i])
			}
		}()

		for v := range ch1 {
			v := v
			for i := 0; i < len(listens); i++ {
				item := i
				if isSync {
					listens[item] <- v
				} else {
					go func() {
						listens[item] <- v
					}()
				}
			}
		}

	}()
}
