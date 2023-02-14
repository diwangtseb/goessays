package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

type Greet struct{}

func (g *Greet) Hello(in1 string, in2 *string) error {
	fmt.Printf("Hello %s %s", in1, *in2)
	return nil
}

func main() {
	g := new(Greet)
	err := rpc.Register(g)
	if err != nil {
		panic(err)
	}
	rpc.HandleHTTP()
	err = http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}
