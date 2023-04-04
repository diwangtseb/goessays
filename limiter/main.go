package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

const bucket_size = 100

var l = rate.NewLimiter(1, bucket_size)

func main() {
	for i := 0; i < 120; i++ {
		// ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		// defer cancel()
		go mockReq(context.TODO())
		time.Sleep(time.Millisecond * 100)
	}
	select {}
}

func mockReq(ctx context.Context) {
	l.Wait(ctx)
	// if err != nil {
	// fmt.Println("token is not ok", rl.Tokens(), err.Error())
	// return
	// }
	// rr := rl.Reserve()
	// dur := rr.Delay()
	// time.Sleep(dur)
	log.Println("token is ok", l.Tokens())
}
