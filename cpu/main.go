package main

import (
	"fmt"
	"runtime"
)

func main() {
	procs := runtime.GOMAXPROCS(2)
	fmt.Println("procs:", procs)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
		runtime.Gosched()
	} // 占用CPU
}
