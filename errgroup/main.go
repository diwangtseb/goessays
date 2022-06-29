package main

import (
	"errors"
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
	return errors.New("gError")
}
