package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			fmt.Printf("%p\n", &close)
			defer wg.Done()
			v, ok := <-close
			fmt.Println(v, ok, num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Printf("timeout exit \n")
	}
	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan struct{})
	tch := time.After(timeout)
	go func(wg *sync.WaitGroup) {
		defer func() {
			ch <- struct{}{}
		}()
		wg.Wait()
	}(wg)
	select {
	case <-ch:
		return false
	case <-tch:
		return true
	}
}
