package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type TokenRecursiveMutex struct {
	sync.Mutex
	token     int64 // 当前持有锁的token
	recursion int32 // 重入的次数
}

// 请求锁需要传入token
func (m *TokenRecursiveMutex) Lock(token int64) {
	// 如果传入的token和持有锁的token一致，此时递归加锁，重入次数+1
	if atomic.LoadInt64(&m.token) == token {
		m.recursion++
		return
	}

	m.Mutex.Lock()
	// 上面的条件不满足，说明不是递归调用，抢到新锁后记录此锁拥有者token值
	atomic.StoreInt64(&m.token, token)
	m.recursion = 1
}

// 释放锁
func (m *TokenRecursiveMutex) Unlock(token int64) {
	// 如果传入token 和持有锁的token不一致，不能解锁，直接运行恐慌
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("wrong the owner(%d):%d", m.token, token))
	}

	// 否则的话把当前此有锁的token重入次数减去1
	m.recursion--

	// 如果发现当前持有锁的token重入次数不等于0，说明还有其他重入锁未解锁
	if m.recursion != 0 {
		return
	}
	// 没有其他重入锁了，释放锁 之前清楚当前锁的持有者token
	atomic.StoreInt64(&m.token, 0)
	// 解锁
	m.Mutex.Unlock()
}

func main() {
	r := &TokenRecursiveMutex{}
	StartTokenLayer(r, 1024)
}

func StartTokenLayer(r *TokenRecursiveMutex, token int64) {
	r.Lock(token)
	fmt.Println("开始")
	TwoTokenLayer(r, token)
	r.Unlock(token)
}

func TwoTokenLayer(r *TokenRecursiveMutex, token int64) {
	r.Lock(token)
	fmt.Println("进入第二层")
	ThreeTokenLayer(r, token)
	r.Unlock(token)
}

func ThreeTokenLayer(r *TokenRecursiveMutex, token int64) {
	r.Lock(token)
	fmt.Println("最后一层")
	r.Unlock(token)
}
