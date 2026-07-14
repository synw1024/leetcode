package main

import (
	"fmt"
	"math"
)

func main() {
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	dp0 := 0
	dp1 := -prices[0]
	dp2 := -math.MaxInt64
	for i := 1; i < l; i++ {
		dp1 = max(dp1, dp0-prices[i])
		dp2 = max(dp2, dp1+prices[i])
	}
	return max(dp0, dp2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
