package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		wg.Add(1)
		go func(i int) {
			fmt.Println("goroutine num is ", i)
			wg.Add(-1) //replace wg.Done()
		}(i)
	}
	wg.Wait()
}
