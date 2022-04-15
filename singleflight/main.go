package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	SingleFlightExample()
}

func SingleFlightExample() {
	var g singleflight.Group
	gsKey := func(reqId int, key string) (interface{}, error) {
		// fmt.Printf("request %v start to get and set cache... \n", reqId)
		value, err, _ := g.Do(key, func() (interface{}, error) {
			fmt.Printf("request %v start to get and set cache... \n", key)
			return "val", nil
		})
		return value.(string), err
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			key := "key"
			_, err := gsKey(i, key)
			if err != nil {
				// handle error
				fmt.Println(err)
			}
			// use value
			// fmt.Println(i, value)
		}(i)
	}
	time.Sleep(time.Second * 20)
}
