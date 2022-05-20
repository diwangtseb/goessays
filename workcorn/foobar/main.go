package main

import (
	"fmt"
	hello "helloworld"
	"time"
)

func main() {
	hello.Hello()
	fmt.Println("end")
	time.Sleep(time.Second * 10)
}
