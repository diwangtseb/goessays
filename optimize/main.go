package main

import (
	"fmt"
	// _ "runtime"
	"time"
	_ "unsafe"
)

//go:linkname nanotime1 runtime.nanotime
func nanotime1() int64
func main() {

	defer func(begin int64) {
		cost := (nanotime1() - begin) / 1000 / 1000
		fmt.Printf("cost = %dms \n", cost)
	}(nanotime1())

	time.Sleep(time.Second)
}
