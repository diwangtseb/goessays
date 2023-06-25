package main

import (
	"fmt"
)

func main() {
	r := minCoin([]int{3, 3, 3}, 10)
	fmt.Println(r)
	r = maxSubArray([]int{1, -2, 4, 3, -1, 0})
	fmt.Println(r)
	r = maxSubSequenceCount([]int{1, 2, 5, 4, 7})
	fmt.Println(r)
	r = stepProblem(3)
	fmt.Println(r)
	r = advanceProblem(4)
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

// max sub sequence count
// 1. 确定问题的最优子结构
// 2. 定义状态空间
// 3. 找到状态转移方程
// 4. 确定边界条件
// 5. 计算状态值
func maxSubSequenceCount(nums []int) int {
	n := len(nums)
	// 边界状态
	if n <= 1 {
		return n
	}
	// 定义状态空间
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	count := make([]int, n)
	for i := 0; i < n; i++ {
		count[i] = 1
	}
	// 最优子结构
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			// 状态方程
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					// 计算状态值
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else {
					if dp[j]+1 == dp[i] {
						count[i] += count[j]
					}
				}
			}
		}
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] >= max {
			max = dp[i]
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if dp[i] == max {
			res += count[i]
		}
	}
	return res
}

// step problem
// case: n = 10, step 1 or 2,there are several ways
// 1. 确定问题的最优子结构
// 2. 定义状态空间
// 3. 找到状态转移方程
// 4. 确定边界条件
// 5. 计算状态值
func stepProblem(n int) (count int) {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	for i := 3; i < n+1; i++ {
		dp[1], dp[2] = 1, 2
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func advanceProblem(n int) (count int) {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}
