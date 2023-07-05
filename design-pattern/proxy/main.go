package main

import "fmt"

func main() {
	ch := &ClientHandle{
		client: NewClienter(),
	}

	ch.DoA("platform")
}

type Clienter interface {
	A()
	B()
}

type client struct{}

// A implements Clienter.
func (*client) A() {
	fmt.Println("执行原始的 A() 方法")
}

// B implements Clienter.
func (*client) B() {
	panic("unimplemented")
}

var _ Clienter = (*client)(nil)

type ClientHandle struct {
	client Clienter
}

func (ch *ClientHandle) DoA(platform string) {
	ch.client.A()
}

type ClientProxy struct {
	client Clienter
}

// B implements Clienter.
func (cp *ClientProxy) B() {
	panic("unimplemented")
}

func (cp *ClientProxy) A() {
	// 在调用原始 client 实例的 A() 方法之前执行拦截逻辑
	fmt.Println("执行拦截逻辑 - 前")

	// 调用原始 client 实例的 A() 方法
	cp.client.A()

	// 在调用原始 client 实例的 A() 方法之后执行拦截逻辑
	fmt.Println("执行拦截逻辑 - 后")
}

func NewClienter() Clienter {
	return &ClientProxy{
		client: &client{},
	}
}
