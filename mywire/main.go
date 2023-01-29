package main

import (
	"fmt"

	"github.com/google/wire"
)

func main() {
	fmt.Println("Hello Wire")
}

type Foo struct{}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct{}

func NewBar() *Bar {
	return &Bar{}
}

type Baz struct {
	b *Bar
	f *Foo
}

func NewBaz(b *Bar, f *Foo) Baz {
	return Baz{
		b: b,
		f: f,
	}
}

// create provider set
var SuperSet = wire.NewSet(NewFoo, NewBar, NewBaz)
