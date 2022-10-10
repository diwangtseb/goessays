package main

import "testing"

func Test_threadOne(t *testing.T) {
	go threadOne()
	go threadTwo()
	select {}
}
