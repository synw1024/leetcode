package main

import "fmt"

func main() {
	res := minPathSum([][]int{[]int{1, 2}, []int{1, 1}})
	fmt.Println(res)
}

// dp[i][j] = min(dp[i-1][j]+dgridp[i][j], dp[i][j-1]+grid[i][j])
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			a, b := -1, -1
			if i-1 >= 0 {
				a = dp[i-1][j]
			}
			if j-1 >= 0 {
				b = dp[i][j-1]
			}
			mm := min(a, b)
			dp[i][j] = mm + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a == -1 && b == -1 {
		return 0
	}
	if a == -1 {
		return b
	}
	if b == -1 {
		return a
	}
	if a < b {
		return a
	}
	return b
}
