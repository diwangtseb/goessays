package main

import "fmt"

type cmder interface {
	Start()
	Stop()
}

type cmdable func()

func (c cmdable) Start() {
	c()
}
func (c cmdable) Stop() {
	c()
}

type client struct {
	cmdable
	hooks
}

func (c *client) Process() {
	c.hooks.processHook()
}

func (c *client) init() {
	c.cmdable = c.Process
	c.hooks.setProcess(func() {
		fmt.Println("base cmd")
	})
}

func NewClient() *client {
	c := client{}
	c.init()
	return &c
}

func main() {
	c := NewClient()
	c.AddHook(&demoHook{})
	c.AddHook(&anotherHook{})
	c.Start()
	c.Stop()
}

type Hook interface {
	ProcessHook(hook ProcessHook) ProcessHook
}

type ProcessHook func()

type hooks struct {
	slice       []Hook
	processHook ProcessHook
}

func (hs *hooks) AddHook(hook Hook) {
	hs.slice = append(hs.slice, hook)
	hs.processHook = hook.ProcessHook(hs.processHook)
}

func (hs *hooks) setProcess(process ProcessHook) {
	hs.processHook = process
	for _, h := range hs.slice {
		if wrapped := h.ProcessHook(hs.processHook); wrapped != nil {
			hs.processHook = wrapped
		}
	}
}

type demoHook struct{}

// ProcessHook implements Hook
func (dh *demoHook) ProcessHook(hook ProcessHook) ProcessHook {
	return func() {
		fmt.Println("hook start")
		defer func() {
			fmt.Println("hook end")
		}()

		hook()
	}
}

var _ Hook = (*demoHook)(nil)

type anotherHook struct{}

// ProcessHook implements Hook
func (*anotherHook) ProcessHook(hook ProcessHook) ProcessHook {
	return func() {
		fmt.Println("another hook start")
		defer func() {
			fmt.Println("another hook end")
		}()

		hook()
	}
}

var _ Hook = (*anotherHook)(nil)
