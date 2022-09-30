package main

import (
	"fmt"
	"net"
)

const First = "Hello World"

func main() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		OnError(err)
	}
	buffer := []byte(First)
	n, err := conn.Write(buffer)
	OnError(err)
	fmt.Println("write", n, string(buffer))
	for {
		recevieMsg(conn)
	}
}

func recevieMsg(conn net.Conn) {
	size := 1024
	buffer := make([]byte, size)
	r, err := conn.Read(buffer)
	OnError(err)
	fmt.Println(r, string(buffer))
}

func OnError(err error) {
	if err != nil {
		panic(err)
	}
}
