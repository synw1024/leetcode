package main

import (
	"fmt"
)

func main() {
	res := rob([]int{1, 2, 3, 1})
	fmt.Println(res)
}

func rob(nums []int) int {
	l := len(nums)

	dp := make([][2]int, l)
	dp[0][1] = nums[0]

	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + nums[i]
	}

	return max(dp[l-1][0], dp[l-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
