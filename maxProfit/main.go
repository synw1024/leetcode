package main

import "fmt"

func main() {
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	l := len(prices)
	dp := make([][2]int, l)
	dp[0][1] = -prices[0]
	for i := 1; i < l; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}
	return dp[l-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
