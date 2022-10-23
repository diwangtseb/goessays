package main

import "fmt"

type temp struct{}

func (t *temp) Add(elem int) *temp {
	fmt.Println(elem)
	return &temp{}
}

func main() {
	tt := &temp{}
	defer tt.Add(1).Add(2)
	tt.Add(3)
}
