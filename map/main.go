package main

import (
	"fmt"
	"time"
)

func main() {
	m := map[string]string{
		"k":  "v",
		"k1": "v",
		"k2": "v",
		"k3": "v",
		"k4": "v",
		"k5": "v",
		"k6": "v",
	}
	for i := 0; i < 10000; i++ {
		go func(x int) {
			updateMap(m)
			fmt.Println(x)
		}(i)
	}
	time.Sleep(20 * time.Second)
}

func updateMap(src map[string]string) map[string]string {
	dst := make(map[string]string)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
