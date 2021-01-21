package main

import (
	"fmt"
	"time"
)

type ticket struct{}

type Mutex struct {
	ch chan ticket
}

// 创建一个缓冲区为1的通道作
func newMutex() *Mutex {
	return &Mutex{ch: make(chan ticket, 1)}
}

// 谁能往缓冲区为1的通道放入数据，谁就获取了锁
func (m *Mutex) Lock() {
	m.ch <- struct{}{}
}

// 解锁就把数据取出
func (m *Mutex) unLock() {
	select {
	case <-m.ch:
	default:
		panic("已经解锁了")
	}
}

func main() {
	mutex := newMutex()
	go func() {
		// 如果是1先拿到锁，那么2就要等1秒才能拿到锁
		mutex.Lock()
		fmt.Println("任务1拿到锁了")
		time.Sleep(1 * time.Second)
		mutex.unLock()
	}()
	go func() {
		mutex.Lock()
		// 如果是2拿先到锁，那么1就要等2秒才能拿到锁
		fmt.Println("任务2拿到锁了")
		time.Sleep(2 * time.Second)
		mutex.unLock()
	}()
	time.Sleep(500 * time.Millisecond)
	// 用了一点小手段这里最后才能拿到锁
	mutex.Lock()
	mutex.unLock()
	close(mutex.ch)
}
