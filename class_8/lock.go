package main

import (
	"net"
	"sync"
	"time"
)

var conMu sync.Mutex
var conn net.Conn

func getConn() net.Conn {
	conMu.Lock()
	defer conMu.Unlock()
	if conn != nil {
		return conn
	}
	conn, _ = net.DialTimeout("TCP", "baidu.com:80", 10*time.Second)
	return conn
}

func main() {
	conn := getConn()
	if conn == nil {
		panic("conn is nil")
	}
}
