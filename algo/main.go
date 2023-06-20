package main

import (
	"fmt"
)

func main() {
	r := minCoin([]int{3, 3, 3}, 10)
	fmt.Println(r)
}

func minCoin(pockets []int, targetCoin int) int {
	dp := make([]int, targetCoin+1)
	for i := 1; i <= targetCoin; i++ {
		dp[i] = targetCoin + 1
	}
	for i := 1; i <= targetCoin; i++ {
		for _, coin := range pockets {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[targetCoin] > targetCoin {
		return -1
	}
	return dp[targetCoin]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
