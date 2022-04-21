package main

import (
	"fmt"
	"strings"
)

func main() {
	slice := []string{"1", "2"}
	r := strings.Join(slice, "','")
	fmt.Printf("'%s'\n", r)
}
