package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
