package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type Once struct {
	done uint32
	m    sync.Mutex
}

func (o *Once) Do(fn func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.doSlow(fn)
}

func (o *Once) doSlow(fn func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = fn()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

func main() {
	urls := []string{
		"127.0.0.1:3453",
		"127.0.0.1:9002",
		"127.0.0.1:9003",
		"baidu.com:80",
	}
	var conn net.Conn
	var o Once
	count := 1
	var err error
	for _, url := range urls {
		err := o.Do(func() error {
			fmt.Printf("初始化%d次\n", count)
			conn, err = net.DialTimeout("tcp", url, time.Second)
			fmt.Println(err)
			return err
		})
		if err == nil {
			break
		}
		count++
		if count == 3 {
			fmt.Println("初始化失败，不再重试")
			break
		}
	}

	if conn != nil {
		_, _ = conn.Write([]byte("GET / HTTP/1.1\r\nHost: google.com\r\n Accept: */*\r\n\r\n"))
		_, _ = io.Copy(os.Stdout, conn)
	}

}
