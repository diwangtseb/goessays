package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithDeadline(ctx, time.Now().Add(time.Second*1))
	go RunRedis(ctx)
	go RunMysql(ctx)
	go RunOther(ctx)
	time.Sleep(time.Second * 10)
}

func RunRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("redis is over", ctx.Err())
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("redis is connecting")
		}
	}
}

func RunMysql(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("mysql is over")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("mysql is connecting")
		}
	}
}

func RunOther(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("other is over")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("other is connecting")
		}

	}
}
