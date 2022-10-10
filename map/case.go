package main

import (
	"fmt"
	"sync"
	"time"
)

var rwLock sync.RWMutex
var muxLock sync.Mutex

func threadOne() {
	fmt.Println("threadOne start")
	muxLock.Lock()
	fmt.Println("threadOne lock")
	time.Sleep(time.Second * 2)
	fmt.Println("threadOne time")
	muxLock.Unlock()
	fmt.Println("threadOne unlock")
}

func threadTwo() {
	fmt.Println("threadTwo start")
	rwLock.RLock()
	fmt.Println("threadTwo lock")
	time.Sleep(time.Second * 2)
	fmt.Println("threadTwo time")
	rwLock.RUnlock()
	fmt.Println("threadTwo unlock")
}
