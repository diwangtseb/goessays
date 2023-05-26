package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	pool := NewGoPool(100)
	count := &atomic.Int32{}
	for i := 0; i < 100; i++ {
		ctx := context.WithValue(context.Background(), DEFAULT_CTX_KEY_XX, tostr(i))
		// ctx, cancel := context.WithTimeout(ctx, time.Second*1)
		go pool.Go(func(ctx context.Context, args ...interface{}) {

			// defer cancel()
			fmt.Println(args)
			if args[0].(int)%2 == 0 {
				fmt.Println("-------------------------", tostr(args[0].(int)))

				ran := 3 + rand.Int31n(5)
				time.Sleep(time.Second * time.Duration(ran))
			}
			count.Add(1)
		}, ctx, i, count)
	}

	// time.Sleep(time.Second * 5)
	// fmt.Println(count)
	// pre := time.Now().Add(time.Second * 5)
	for {
		if len(pool.semaphore) == 0 {
			fmt.Println("return pre当前线程池数量----", len(pool.semaphore))
			return
		}
		time.Sleep(time.Millisecond * 500)
		fmt.Println("current 当前线程池数量----", len(pool.semaphore))
		// if time.Now().After(pre) {
		// 	fmt.Println("break 当前线程池数量----", len(pool.semaphore))
		// 	break
		// }
	}
}

type Key string

const DEFAULT_CTX_KEY_XX Key = Key("xx")

type GoPool struct {
	pool        *sync.Pool
	semaphore   chan struct{}
	maxRoutines int
}

func NewGoPool(maxRNum int) *GoPool {
	return &GoPool{
		pool: &sync.Pool{New: func() interface{} {
			return make(chan struct{}, 1)
		}},
		semaphore:   make(chan struct{}, maxRNum),
		maxRoutines: maxRNum,
	}
}

func (gp *GoPool) worker(ctx context.Context, f func(context.Context, ...interface{}), args ...interface{}) {
	ch := gp.pool.Get().(chan struct{})
	defer func() {
		<-gp.semaphore
		<-ch
		gp.pool.Put(ch)
		if err := recover(); err != nil {
			fmt.Println("<GoPool> panic", err)
		}
	}()

	f(ctx, args...)
	select {
	case <-ctx.Done():
		var errstr string
		v := ctx.Value(DEFAULT_CTX_KEY_XX)
		if ctxVal, ok := v.(string); ok {
			errstr += ctxVal
		}
		outErr := fmt.Errorf("%v: %v", ctx.Err(), errstr)
		panic(outErr)
	default:
		ch <- struct{}{}
		statCurrentRoutineNum(len(gp.semaphore))
	}
}
func statCurrentRoutineNum(num int) {
	fmt.Printf("current run goroutine num is %d\n", num)
}

func (gp *GoPool) Go(f func(context.Context, ...interface{}), ctx context.Context, args ...interface{}) {
	gp.semaphore <- struct{}{}
	go gp.worker(ctx, f, args...)
}

func tostr[T int | uint](t T) string {
	return fmt.Sprintf("%d", t)
}
