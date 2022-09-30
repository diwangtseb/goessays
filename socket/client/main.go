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
}

func OnError(err error) {
	if err != nil {
		panic(err)
	}
}
