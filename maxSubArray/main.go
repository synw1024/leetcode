package main

import (
	"fmt"
)

func main() {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}

// func maxSubArray(nums []int) int {
// 	l := len(nums)
// 	dp := make([][2]int, l)
// 	dp[0][0] = 0
// 	dp[0][1] = nums[0]
// 	for i := 1; i < l; i++ {
// 		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
// 		dp[i][1] = max(dp[i-1][0]+nums[i], dp[i-1][1]+nums[i])
// 	}
// 	return max(dp[l-1][0], dp[l-1][1])
// }

func maxSubArray(nums []int) int {
	l := len(nums)
	f := make([]int, l)
	if nums[0] > 0 {
		f[0] = nums[0]
	}
	for i := 1; i < l; i++ {
		f[i] = max(f[i-1], f[i-1]+nums[i])
	}
	return f[l-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// f[i] = max{f[i-1], f[i-1]+arr[i]}
