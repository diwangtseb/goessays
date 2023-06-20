package main

import (
	"fmt"
)

func main() {
	r := minCoin([]int{3, 3, 3}, 10)
	fmt.Println(r)
	r = maxSubArray([]int{1, -2, 4, 3, -1, 0})
	fmt.Println(r)
}

// min coins equal target coin in pockets
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

// max continuous sum sub array
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	ans := dp[0]
	for i := 1; i < n; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
