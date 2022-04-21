package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go Send()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Receive()
		time.Sleep(time.Second * 5)
	}
	wg.Wait()
}

var ch = make(chan int, 100)

func Send() {
	for i := 0; i < 1000; i++ {
		ch <- i
		fmt.Println("Sent", i)
	}
}

func Receive() []int {
	var r []int
	var count int
	for {
		select {
		case condition := <-ch:
			fmt.Println("Received", condition)
			r = append(r, condition)
			count++
			if count > 100 {
				fmt.Println("Received 100 messages")
				return r
			}
		default:
			return r
		}
		time.Sleep(time.Millisecond * 100)
	}
}
