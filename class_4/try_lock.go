package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	mutexLocked      = 1 << iota // 加锁标示位置
	mutexWoken                   // 唤醒标示位置
	mutexStarving                // 锁饥饿标示
	mutexWaiterShift = iota      // 标示Waiter的起始bit位置
)

type Mutex struct {
	sync.Mutex
}

// 尝试获取锁
func (m *Mutex) TryLock() bool {
	// fast path 如果一开始就没有其他g争夺，那么直接获取锁
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}
	// 如果处于唤醒，加锁或者饥饿状态，这次请求就不参与了竞争了，直接返回false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
		return false
	}
	// 尝试在竞争的状态下请求锁
	n := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, n)
}

func main() {
	var mu Mutex
	go func() { // 启动一个g持有一段时间的锁
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)
	ok := mu.TryLock() // 尝试获取锁
	if ok {            // 获取锁成功
		fmt.Println("got the lock")
		// 开始你的业务
		mu.Unlock()
		return
	}
	// 没有获取到
	fmt.Println("can't get the lock")
}
