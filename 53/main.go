package main

import (
	"fmt"
	"math"
)

func main() {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	l := len(nums)

	dp := make([][3]int, l)
	dp[0][1] = nums[0]
	dp[0][2] = -math.MaxInt64

	for i := 1; i < l; i++ {
		dp[i][1] = max(dp[i-1][0]+nums[i], dp[i-1][1]+nums[i])
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}

	return max(dp[l-1][1], dp[l-1][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
