package main

import (
	"fmt"
)

func main() {
	m := make(map[User]struct{})
	u1 := User{Id: "cc", Name: "cc"}
	m[u1] = struct{}{}

	u2 := User{Id: "cc2"}
	m[u2] = struct{}{}

	u3 := User{Id: "cc3"}
	_, ok2 := m[u2]
	_, ok3 := m[u3]
	fmt.Println(ok2, ok3)
}

type User struct {
	Id   string
	Name string
	Last int
}
