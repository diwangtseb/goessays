package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
)

const bucket_size = 10

func main() {
	l := rate.NewLimiter(1, bucket_size)
	for i := 0; i < 20; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		go mockReq(ctx, l)
		time.Sleep(time.Millisecond * 100)
	}
	select {}
}

func mockReq(ctx context.Context, rl *rate.Limiter) {
	err := rl.Wait(ctx)
	if err != nil {
		fmt.Println("token is not ok", rl.Tokens(), err.Error())
		return
	}
	// rr := rl.Reserve()
	// dur := rr.Delay()
	log.Println("token is ok", rl.Tokens())
}
