package main

import (
	"fmt"
	"math"
)

func main() {
	res := jump([]int{2, 3, 1, 1, 4})
	fmt.Println(res)
}

// dp[i] = min(dp[i-step]+1)
func jump(nums []int) int {
	l := len(nums)
	dp := make([]int, l)

	for i := 1; i < l; i++ {
		min := math.MaxInt64
		for j := i - 1; j >= 0; j-- {
			if nums[j] < i-j {
				continue
			}
			if dp[j] < min {
				min = dp[j]
			}
		}
		dp[i] = min + 1
	}
	return dp[l-1]
}
