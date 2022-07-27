package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group
	for i := 0; i < 10; i++ {
		eg.Go(gError)
	}
	err := eg.Wait()
	fmt.Println(err)
}

func gError() error {
	fmt.Println("run here")
	return nil
}
