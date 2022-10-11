package main

import (
	"fmt"
	"testing"
	"time"
)

func Test_getAttrs(t *testing.T) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		attrs := getAttrs(key)
		go func() {
			attrs = getAttrs(key)
			fmt.Println(attrs)
		}()

		go func() {
			setAttrs(key, attrs)
		}()
	}
	time.Sleep(10 * time.Second)
}
