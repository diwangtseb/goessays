package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b, c := a[:3], a[3:]
	b[2] = 100
	fmt.Println(a, b, c)
}
