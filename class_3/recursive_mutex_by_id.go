package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// 简单获取gid
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
//	gid:=GoID()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}



func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	//	gid:=GoID()

	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner (%d):%d!", m.owner, gid))
	}
	m.recursion--
	if m.recursion != 0 { // 如果这个g还没完全释放，则直接返回
		return
	}
	// 此g最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

// 通过runtime.Stack 获取运行gID
func GoId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get g id:%v", err))
	}
	return id
}

func main() {
	r := &RecursiveMutex{}
	StartLayer(r)
}

func StartLayer(r *RecursiveMutex) {
	r.Lock()
	fmt.Println("开始")
	TwoLayer(r)
	r.Unlock()
}

func TwoLayer(r *RecursiveMutex) {
	r.Lock()
	fmt.Println("进入第二层")
	ThreeLayer(r)
	r.Unlock()
}

func ThreeLayer(r *RecursiveMutex) {
	r.Lock()
	fmt.Println("最后一层")
	r.Unlock()
}
