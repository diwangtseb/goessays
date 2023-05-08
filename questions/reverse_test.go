package main

import (
	"fmt"
	"testing"
)

func StrReverse(src []byte) {
	mid := int(len(src) / 2)
	for i := 0; i < len(src); i++ {
		if i == mid {
			break
		}
		src[i], src[len(src)-1-i] = src[len(src)-1-i], src[i]
	}
}

func TestReverse(t *testing.T) {
	src := []byte("")
	fmt.Println(string(src))
	StrReverse(src)
	fmt.Println(string(src))
}
