package main

import (
	"fmt"
	"time"
)

func main() {
	p := NewPublisher(10, time.Second*1)
	foo := p.Subscribe("foo", func(topic interface{}) bool {
		return topic.(string) == "foo"
	})
	all := p.Subscribe("all", nil)
	p.Publish("foo")
	go func() {
		for c := range foo {
			fmt.Println(c)
		}
	}()

	go func() {
		for a := range all {
			fmt.Println(a, "all")
		}
	}()
	p.Publish("bar")
	time.Sleep(time.Second * 5)
}

type (
	topicFunc func(interface{}) bool
)

type Publisher struct {
	buffer      int
	timeOut     time.Duration
	subscribers map[chan interface{}]topicFunc
}

func NewPublisher(buffer int, timeOut time.Duration) Publisher {
	return Publisher{
		buffer:      buffer,
		timeOut:     timeOut,
		subscribers: make(map[chan interface{}]topicFunc),
	}
}

func (p *Publisher) Publish(topic interface{}) {
	for ch, f := range p.subscribers {
		go func(ch chan interface{}, f topicFunc, msg interface{}) {
			if f != nil && !f(topic) {
				return
			}
			ch <- topic
		}(ch, f, topic)
	}
}

func (p *Publisher) Subscribe(topic interface{}, f topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.subscribers[ch] = f
	return ch
}
