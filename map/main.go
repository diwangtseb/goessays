package main

import (
	"fmt"
	"sync"
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
	var mux sync.RWMutex
	for i := 0; i < 100; i++ {
		go func() {
			mux.Lock()
			m["k"] = "ooo"
			m = updateMap(m)
			mux.Unlock()
		}()
	}
	mux.RLock()
	fmt.Println("pre", m)
	mux.RUnlock()
	time.Sleep(2 * time.Second)
	fmt.Println("later", m)
}

func updateMap(src map[string]string) map[string]string {
	dst := make(map[string]string)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
