package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var g Group
	g.GOMAXPROCS(1) // 只使用一个goroutine处理子任务

	var count int64
	g.Go(func(ctx context.Context) error {
		time.Sleep(time.Second) // 睡眠5秒，把这个goroutine占住
		return nil
	})

	total := 100

	for i := 0; i < total; i++ { // 并发一万个goroutine执行子任务，理论上这些子任务都会加入到Group的待处理列表中
		go func() {
			g.Go(func(ctx context.Context) error {
				atomic.AddInt64(&count, 1)
				return nil
			})
		}()
	}

	// 等待所有的子任务完成。理论上101个子任务都会被完成
	if err := g.Wait(); err != nil {
		panic(err)
	}

	got := atomic.LoadInt64(&count)
	if got != int64(total) {
		panic(fmt.Sprintf("expect %d but got %d", total, got))
	}
}



// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	err     error
	wg      sync.WaitGroup
	errOnce sync.Once

	workerOnce sync.Once
	ch         chan func(ctx context.Context) error
	chs        []func(ctx context.Context) error

	ctx    context.Context
	cancel func()
	mu sync.Mutex
}

func (g *Group) do(f func(ctx context.Context) error) {
	ctx := g.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	var err error
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 64<<10)
			buf = buf[:runtime.Stack(buf, false)]
			err = fmt.Errorf("errgroup: panic recovered: %s\n%s", r, buf)
		}
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
		g.wg.Done()
	}()
	err = f(ctx)
}

// GOMAXPROCS set max goroutine to work.
func (g *Group) GOMAXPROCS(n int) {
	if n <= 0 {
		panic("errgroup: GOMAXPROCS must great than 0")
	}
	g.workerOnce.Do(func() {
		g.ch = make(chan func(context.Context) error, n)
		for i := 0; i < n; i++ {
			go func() {
				for f := range g.ch {
					g.do(f)
				}
			}()
		}
	})
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Group) Go(f func(ctx context.Context) error) {
	g.wg.Add(1)
	if g.ch != nil {
		select {
		case g.ch <- f:
		default:
			g.chs = append(g.chs, f)
		}
		return
	}
	go g.do(f)
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	if g.ch != nil {
		for _, f := range g.chs {
			g.ch <- f
		}
	}
	g.wg.Wait()
	if g.ch != nil {
		close(g.ch) // let all receiver exit
	}
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
