package main

import "fmt"

type User struct {
	Name string
	M    map[string]string
}

func main() {
	a := &User{
		Name: "a",
		M: map[string]string{
			"a": "a",
		},
	}
	b := *a
	fmt.Printf("%p,%p", &a.M, &b.M)
}
