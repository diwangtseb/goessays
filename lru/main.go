package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
)

func main() {
	LRUExample()
}

func LRUExample() {
	l, _ := lru.New(128)
	for i := 0; i < 256; i++ {
		l.Add(i, nil)
	}
	if l.Len() != 128 {
		panic(fmt.Sprintf("bad len: %v", l.Len()))
	}
	fmt.Println(l.Keys()...)
}
