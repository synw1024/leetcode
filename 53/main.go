package main

import (
	"fmt"
)

func main() {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}

func maxSubArray(nums []int) (res int) {
	l := len(nums)

	dp := make([]int, l)
	dp[0] = nums[0]

	res = dp[0]
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
