package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	SleepSort(1, 4, 2, 3, 5, 7, 6, 8, 9, 10)
	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		default:
			continue
		}
	}
}

func SleepSort(nums ...int) {
	for _, num := range nums {
		go func(num int) {
			time.Sleep(time.Duration(num) * time.Nanosecond)
			ch <- num
		}(num)
	}
}
