package main

import (
	"fmt"
	"math"
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
	if l < 2 {
		return nums[0]
	}
	a := 0
	b := max(a, nums[0])
	c := -math.MaxInt64
	for i := 1; i < l; i++ {
		b = max(a, b+nums[i])
		c = max(b, c)
	}
	if c != 0 {
		return c
	} else {
		max := -math.MaxInt64
		for i := 0; i < l; i++ {
			if max < nums[i] {
				max = nums[i]
			}
		}
		return max
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
