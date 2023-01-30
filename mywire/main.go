package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/google/wire"
)

func main() {
	app, clear := initialize(context.Background())
	defer clear()
	app.Run()
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

func NewBaz(b *Bar, f *Foo) *Baz {
	return &Baz{
		b: b,
		f: f,
	}
}

type Fooer interface {
	Run()
}

type MyFooer string

// Foo implements Fooer
func (mf MyFooer) Run() {
	fmt.Println("run ...")
}

func NewMyfoo(baz *Baz, fc *FooCase, vfc *ValueFooCase, itf Fooer, itfIo io.Reader, ffc string) (MyFooer, func()) {
	a := "hello my fooer"
	fmt.Printf("%+v", baz)
	fmt.Printf("%+v", fc)
	fmt.Printf("%+v", vfc)
	fmt.Printf("%+v", itf)
	fmt.Printf("%+v", itfIo)
	fmt.Printf("%+v", ffc)
	mf := MyFooer(a)
	clearUp := func() {
		fmt.Println("clear ...")
	}
	return mf, clearUp
}

var _ Fooer = (*MyFooer)(nil)

// var Set = wire.Bind(new(Fooer), new(MyFooer))

// create provider set
var SuperSet = wire.NewSet(NewFoo, NewBar, NewBaz, NewMyfoo, NewA, NewB, StructSet, ValuesSet, ItfSet, ItfIoSet, FiledSet)

// sturcts
type FooCase struct {
	A string `wire:"-"`
	B int
}

func NewA() string {
	return "a"
}

func NewB() int {
	return 0
}

var StructSet = wire.Struct(new(FooCase), "*")

type ValueFooCase struct {
	A string
	B int
}

var ValuesSet = wire.Value(&ValueFooCase{
	A: "ooooo",
	B: 0,
})

var ItfSet = wire.InterfaceValue(new(Fooer), new(MyFooer))
var ItfIoSet = wire.InterfaceValue(new(io.Reader), os.Stdin)

type FieldsFooCase struct {
	A bool
}

// func NewFieldsFooCase(a bool) FieldsFooCase {
// 	return FieldsFooCase{
// 		A: a,
// 	}
// }

var FiledSet = wire.FieldsOf(new(FieldsFooCase), "A")

// clear partial
