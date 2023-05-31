package main

import (
	"fmt"
	"runtime"
)

type MyObject struct {
	Name string
}

func finalize(o *MyObject) {
	fmt.Printf("Cleaning up %s\n", o.Name)
}

func main() {
	obj := &MyObject{Name: "myobject"}
	runtime.SetFinalizer(obj, finalize)

	// 制造一些垃圾
	for i := 0; i < 10; i++ {
		_ = make([]byte, 10*1024*1024)
	}
}
