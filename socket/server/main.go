package main

import (
	"fmt"
	"net"
	"sync/atomic"
)

const MAX_CONN_NUM = iota + 2

func main() {
	listener, err := net.Listen("tcp", ":9999")
	connChan := make(chan net.Conn, MAX_CONN_NUM)
	OnError(err)
	var cn uint32
	for {
		conn, err := listener.Accept()
		OnError(err)
		v := atomic.AddUint32(&cn, 1)
		if v < MAX_CONN_NUM {
			connChan <- conn
		} else {
			hanldConn(connChan)
		}
	}
}

func hanldConn(ch <-chan net.Conn) {
	select {
	case c := <-ch:
		buffer := make([]byte, 1024)
		n, err := c.Read(buffer)
		OnError(err)
		fmt.Println("read", n, string(buffer))
	default:
		fmt.Println("deafult")
	}
}

func OnError(err error) {
	if err != nil {
		panic(err)
	}
}
