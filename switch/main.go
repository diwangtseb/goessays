package main

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("1")
		fallthrough
	case false:
		fmt.Println("2")
		fallthrough
	case false:
		fmt.Println("3")
	default:
		fmt.Println("luck")
	}
}
