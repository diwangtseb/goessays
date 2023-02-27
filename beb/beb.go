package main

import (
	"math"
	"math/rand"
	"time"
)

func main() {
	maxRetries := 10
	retryCount := 0
	baseTime := 35

	for {
		if !mockReq() {
			time.Sleep(time.Duration(int64(baseTime) * int64(math.Pow(2, float64(retryCount)))))
			retryCount++
		} else {
			break
		}
		if retryCount == maxRetries {
			break
		}
	}
}

var count = 0

func mockReq() bool {
	count = int(base1Rand())
	return count < 5
}

func base1Rand() int64 {
	return int64(rand.Intn(100) + 1)
}
