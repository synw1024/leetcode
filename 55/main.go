package main

import "fmt"

func main() {
	res := canJump([]int{0, 2, 3})
	fmt.Println(res)
}

func canJump(nums []int) bool {
	l := len(nums)
	if l == 1 {
		return true
	}

	dp := make([]bool, l)

	for i := 1; i < l; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
		if !dp[i] {
			return false
		}
	}
	return dp[l-1]
}
