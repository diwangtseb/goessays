package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//mockCtxCancel(context.TODO())
	ctx := mockCtxTimeout(context.TODO())
	v := <-ctx.Done()
	fmt.Println(v)
}

func mockCtxCancel(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println("do something")
	return ctx
}

func mockCtxTimeout(ctx context.Context) context.Context {
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	time.Sleep(time.Second * 2)
	fmt.Println("do something")
	return ctx
}
