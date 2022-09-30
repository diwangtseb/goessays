package main

import (
	"fmt"
	"net"
)

const MAX_CONN_NUM = iota + 2

func main() {
	listener, err := net.Listen("tcp", ":9999")
	connChan := make(chan net.Conn, MAX_CONN_NUM)
	OnError(err)
	for {
		conn, err := listener.Accept()
		OnError(err)
		connChan <- conn
		go hanldConn(connChan)
	}
}

func hanldConn(ch <-chan net.Conn) {
	select {
	case c := <-ch:
		buffer := make([]byte, 1024)
		n, err := c.Read(buffer)
		OnError(err)
		fmt.Println("read", n, string(buffer))
		writeMsg(c)
	default:
		fmt.Println("deafult")
	}
}

func writeMsg(conn net.Conn) {
	msg := []byte("thanks")
	n, err := conn.Write(msg)
	OnError(err)
	fmt.Println("write", n, string(msg))
}

func OnError(err error) {
	if err != nil {
		panic(err)
	}
}
