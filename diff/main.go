package main

import (
	"fmt"

	"k8s.io/utils/diff"
)

func main() {
	a := map[string]string{"a": "1"}
	b := map[string]string{"a": "1"}
	r := diff.ObjectDiff(a, b)
	fmt.Println(r)
}
