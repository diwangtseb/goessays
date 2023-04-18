package main

import (
	"fmt"
	"time"
)

func main() {
	buffer := 10
	ch := make(chan struct{}, buffer)
	for i := 0; i < buffer; i++ {
		go func(i int) {
			order := NewOrder(i, ch)
			order.Start()
		}(i)
	}
	for {
		select {
		case <-ch:
			fmt.Println("ok")
		}
	}
}

type Dosomethinger interface {
	Buy()
	CallBack()
}

type StatufulMachiner interface {
	Start()
	Stop()
	CirCle()
}

type order struct {
	channel chan<- struct{}
	id      int
	status  int
}

func NewOrder(id int, ch chan<- struct{}) *order {
	return &order{
		channel: ch,
		id:      id,
		status:  0,
	}
}

type EnumStatus int

const (
	Start  EnumStatus = 1
	CirCle EnumStatus = 2
	End    EnumStatus = 3
)

// CirCle implements StatufulMachiner
func (o *order) CirCle() {
	o.status = int(CirCle)
	o.Buy()
}

// Start implements StatufulMachiner
func (o *order) Start() {
	o.status = int(Start)
	o.CirCle()
}

// Stop implements StatufulMachiner
func (o *order) Stop() {
	o.status = int(CirCle)
	fmt.Printf("id %d is stop \n", o.id)
	o.channel <- struct{}{}
}

func (o *order) Buy() {
	time.Sleep(time.Second * 2)
	fmt.Printf("id %d buy ok \n", o.id)
	o.CallBack()
}

func (o *order) CallBack() {
	fmt.Printf("id %d buy's callback", o.id)
	o.Stop()
}

var _ StatufulMachiner = (*order)(nil)
