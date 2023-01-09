package main

import "fmt"

func main() {
	func() {
		c := make(chan *Mem, 1)
		// go func() {
		c <- &Mem{
			id: "1",
		}
		// }()
		getCh(c)
	}()
}

type Mem struct {
	id string
}

func getCh(c <-chan *Mem) {
	v := <-c
	fmt.Println(v)
}
