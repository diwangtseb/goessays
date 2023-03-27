package main

import (
	"fmt"
)

func main() {
	if fanInCh == nil {
		fanInCh = make(chan chan *Msg, 100000)
	}
	go produer()
	go distribute()
	select {}
}

type Msg struct {
	req interface{}
	rsp interface{}
}

var fanInCh chan chan *Msg

func produer() {
	for i := 0; i < 100; i++ {
		msgCh := make(chan *Msg, 1)
		msgCh <- &Msg{
			req: i,
			rsp: i,
		}
		fanInCh <- msgCh
	}
}

func distribute() {
	// req := <-proch

	for c := range fanInCh {
		go worker(c)
	}
}

func worker(r <-chan *Msg) {
	msg := <-r
	//do worker
	fmt.Printf("do worker req %d  <=>", msg.req)
	defer func() {
		fmt.Printf("do work rsp %d \n", msg.rsp)
	}()
	msg.rsp = 2222
}
