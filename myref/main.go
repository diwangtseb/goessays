package main

import (
	"fmt"
	"reflect"
)

type A struct {
	b string
}

func (a *A) Hw() {
	fmt.Println("hello :", a.b)
}

func (a *A) HwWithArgs(args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("lens error")
	}
	fmt.Println(args)
	return nil
}

func main() {
	a := A{
		b: "b",
	}
	fa := reflect.ValueOf(&a)
	infoFunc := fa.MethodByName("Hw")
	fmt.Printf("%v", infoFunc)
	infoFunc.Call([]reflect.Value{})
	f2 := fa.MethodByName("HwWithArgs")
	a1 := reflect.ValueOf("zhangsan")
	a2 := reflect.ValueOf("lisi")
	f2.Call([]reflect.Value{a1, a2})
}
