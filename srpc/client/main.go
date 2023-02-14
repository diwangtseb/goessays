package main

import "net/rpc"

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	rsp := new(string)
	err = client.Call("Greet.Hello", "hello", rsp)
	if err != nil {
		panic(err)
	}
}
