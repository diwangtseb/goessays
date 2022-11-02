package main

import (
	"fmt"
	"time"
)

func main() {
	foo := map[string]*Foo{
		"id": {
			Attrs: map[string]string{},
		},
	}
	refMap(foo)
	time.Sleep(time.Second * 5)
	bar := foo
	fmt.Printf("%p,%p", foo, bar)
}

type Foo struct {
	Attrs map[string]string
}

var ch = make(chan interface{}, 1000)

func refMap(src map[string]*Foo) {
	select {
	case ch <- nil:
		fmt.Println("run here")
		go func(s map[string]*Foo) {
			attrs := make(map[string]string)
			if v, ok := src["id"]; ok {
				attrs = v.Attrs
			}
			attrs["k"] = "xx"
		}(src)
	default:
		fmt.Println("---- ")
	}
}
