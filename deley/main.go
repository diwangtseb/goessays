package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dq := &dele_queue{}
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			dq.display()
		}
	}()
	rand.Seed(time.Now().UnixMicro())
	for {
		time.Sleep(time.Second * 1)
		r := rand.Int31n(100)
		rs := fmt.Sprintf("%d", r)
		dq.delay_push(rs)
	}
}

type dele_queue []string

func (dq *dele_queue) delay_push(a string) {
	*dq = append(*dq, a)
}

func (dq *dele_queue) display() {
	fmt.Println("dq ...", dq)
}
