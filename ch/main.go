package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			c := make(chan *Mem, 1)
			defer close(c)
			// go func() {
			c <- &Mem{
				id: "1",
			}
			// }()
			getCh(c)
		}()
	}
	select {}
}

type Mem struct {
	id string
}

func getCh(c <-chan *Mem) {
	v := <-c
	fmt.Println(v)
}
