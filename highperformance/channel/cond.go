package channel

import (
	"fmt"
	"sync"
	"time"
)

type IOMocker interface {
	Read()
	Write()
}

type InnerChan struct {
	ch   chan struct{}
	cond sync.Cond
	flag bool
}

func NewInnerChan() *InnerChan {
	ch := make(chan struct{}, 3)
	return &InnerChan{ch: ch, cond: sync.Cond{L: &sync.Mutex{}}, flag: false}
}

func (ic *InnerChan) Read() {
	fmt.Println("read start")
	ic.cond.L.Lock()
	for !ic.flag {
		ic.cond.Wait()
	}
	fmt.Printf("read %s\n", <-ic.ch)
	ic.cond.L.Unlock()
}

func (ic *InnerChan) Write() {
	fmt.Println("read end")
	ic.ch <- struct{}{}
	ic.ch <- struct{}{}
	ic.ch <- struct{}{}
	time.Sleep(time.Second * 1)
	ic.cond.L.Lock()
	ic.flag = true
	ic.cond.L.Unlock()
	ic.cond.Broadcast()
}
