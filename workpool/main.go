package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	pool := NewGoPool(1)
	count := &atomic.Int32{}
	for i := 0; i < 1000; i++ {
		ctx := context.WithValue(context.Background(), "k->", i)
		// ctx, cancel := context.WithTimeout(ctx, time.Second*1)
		go pool.Go(func(ctx context.Context, args ...interface{}) {
			// defer cancel()
			fmt.Println(args)
			if args[0].(int) == 100 {
				fmt.Println("-------------------------", i)
				f := func(ctx context.Context) {
					time.Sleep(time.Second * 2)
					fmt.Println(ctx.Err())
				}
				f(ctx)
			}
			count.Add(1)
		}, ctx, i, count)
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

func (gp *GoPool) Go(f func(ctx context.Context, args ...interface{}), ctx context.Context, args ...interface{}) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	gp.semaphore <- struct{}{}
	ch := gp.pool.Get().(chan bool) // 获取一个可用的goroutine
	go func() {
		defer cancel()
		defer func() {
			gp.pool.Put(ch) // 将不再使用的goroutine放回对象池中
			if err := recover(); err != nil {
				fmt.Println("panic ->", err)
			}
		}()

		f(ctx, args...) // 执行任务
		for {
			select {
			case <-ctx.Done():
				v := ctx.Value("k->")
				r := tostr(v.(int))
				inErr := errors.New("goroutine num -> " + r + "\n")
				outErr := errors.Join(ctx.Err(), inErr)
				panic(outErr)
			default:
				ch <- true
				return
			}
		}
	}()
	<-gp.semaphore
}

func tostr[T int | uint](t T) string {
	return fmt.Sprintf("%d", t)
}
