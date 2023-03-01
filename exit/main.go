package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	start()
	go close()
	select {}
}

func start() {
	go func() {
		for {
			fmt.Println("receive")
			time.Sleep(time.Second * 1)
		}
	}()
}

func close() {
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, os.Interrupt, os.Kill)
	go func() {
		<-sch
		fmt.Println("secure exit")
	}()
}
