package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type RedPkg struct {
	redID       int
	sumMoney    int
	remainMoney int
	count       int
	lock        sync.Mutex
}

var addr int32 = 0

func NewRedPkg(money, count int) *RedPkg {
	//split red pkg
	rp := RedPkg{
		redID:       int(atomic.AddInt32(&addr, 1)),
		remainMoney: money,
		sumMoney:    money,
		count:       count,
		lock:        sync.Mutex{},
	}
	return &rp
}

func (rp *RedPkg) distributionMoney() {
	rp.lock.Lock()
	defer rp.lock.Unlock()
	if rp.count == 0 {
		fmt.Printf("red pkg %d is empty\n", rp.redID)
		return
	}
	if rp.count == 1 {
		money := rp.remainMoney
		rp.remainMoney -= money
		rp.count--
		fmt.Printf("red pkg %d, get money %d, remain %d\n", rp.redID, money, rp.remainMoney)
		return
	}
	money := RandomMoney(rp.remainMoney, rp.count)
	rp.remainMoney -= money
	rp.count--
	fmt.Printf("red pkg %d, get money %d, remain %d\n", rp.redID, money, rp.remainMoney)
}

func RandomMoney(money, count int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := money / count * 2
	return rand.Intn(max) + min
}

func main() {
	nrp := NewRedPkg(100, 10)
	for i := 0; i < 20; i++ {
		go nrp.distributionMoney()
	}
	time.Sleep(time.Second * 2)
}
