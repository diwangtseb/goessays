package main

import "fmt"

func main() {
	for {
		switch 1 {
		case 0:
			fmt.Println(0)
		case 1:
			fmt.Println(1)
			break
		default:
			fmt.Println("default")
		}
	}
}
