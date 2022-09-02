package main

import (
	"fmt"
	"unsafe"
)

type SelfMap map[string]string

func (s *SelfMap) Error() string {
	r := ""
	for k, v := range *s {
		r += fmt.Sprintf("k{%s},v{%s}", k, v)
	}
	return r
}

type A struct {
	ID int `json:"id"`
}

func (a *A) Error() string {
	return string(rune(a.ID))
}

func main() {
	var sm SelfMap
	fmt.Println(unsafe.Sizeof(sm))
	r := sm.Error()
	fmt.Println("rrr", r)

	var a *A
	r2 := a.Error()
	fmt.Println("r2r2r2", r2)
}
