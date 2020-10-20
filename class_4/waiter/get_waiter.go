package main

import (
	"fmt"
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

func main() {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("waitings:%d,isLocked:%t,woken:%t,starving:%t\n", mu.Count(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}

func (m *Mutex) Count() int {
	// 获取state字段的值
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift // 得到等待者的数
	v = v + (v & mutexLocked) // 再加上锁持有者的数量 0或者1
	return int(v)
}

// 锁是否被持有
func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

// 是否有等待者被唤醒
func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

// 锁是否处于饥饿状态
func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
}
