package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	pool := NewGoPool(1)
	count := &atomic.Int32{}
	for i := 0; i < 1000; i++ {
		go pool.Go(func(args ...interface{}) {
			fmt.Println(args)
			count.Add(1)
		}, i, count)
	}

	time.Sleep(time.Second * 5)
	fmt.Println(count)
}

type GoPool struct {
	pool        *sync.Pool
	semaphore   chan struct{}
	maxRoutines int
}

func NewGoPool(maxRNum int) *GoPool {
	return &GoPool{
		pool: &sync.Pool{New: func() interface{} {
			ch := make(chan bool)
			go func() {
				defer close(ch)
				for {
					select {
					case <-ch:
						return
					default:
						time.Sleep(time.Millisecond * 100)
					}
				}
			}()
			return ch
		}},
		semaphore:   make(chan struct{}, maxRNum),
		maxRoutines: maxRNum,
	}
}

func (gp *GoPool) Go(f func(args ...interface{}), args ...interface{}) {
	gp.semaphore <- struct{}{}
	ch := gp.pool.Get().(chan bool) // 获取一个可用的goroutine
	go func() {
		defer func() {
			gp.pool.Put(ch) // 将不再使用的goroutine放回对象池中
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		f(args...) // 执行任务
		ch <- true
	}()
	<-gp.semaphore
}
