package main

import "time"

func main() {
	a := make(chan struct{})
	go func() {
		time.Sleep(time.Second * 1)
		close(a)
	}()
	go func() {
		time.Sleep(time.Second * 2)
		a <- struct{}{}
	}()
	time.Sleep(time.Second * 3)
}
