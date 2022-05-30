package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

var GLOBALStr []string

func task() {
	for range time.Tick(time.Second * 1) {
		GLOBALStr = append(GLOBALStr, "hello")
	}
}

var g singleflight.Group

func fakeReq() interface{} {
	r, _, _ := g.Do("task", func() (interface{}, error) {
		return GLOBALStr, nil
	})
	return r
}

func main() {
	go task()
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println(fakeReq())
		}()
	}
	time.Sleep(time.Second * 5)
}
