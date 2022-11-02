package main

import "time"

func main() {
	foo := map[string]*Foo{
		"id": {
			Attrs: map[string]string{},
		},
	}
	for i := 0; i < 1000; i++ {
		go refMap(foo)
	}
	time.Sleep(time.Second * 5)
}

type Foo struct {
	Attrs map[string]string
}

func refMap(src map[string]*Foo) {
	for i := 0; i < 10; i++ {
		attrs := make(map[string]string)
		if v, ok := src["id"]; ok {
			attrs = v.Attrs
		}
		attrs["k"] = "xx"
	}

}
