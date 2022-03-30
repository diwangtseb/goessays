package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		mustPanic()
	}()
	time.Sleep(time.Second * 10)
}

func mustPanic() {
	panic("thanks i'm panic")
}
