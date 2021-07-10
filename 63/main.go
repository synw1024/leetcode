package main

import "fmt"

func main() {
	res := uniquePathsWithObstacles([][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}})
	fmt.Println(res)
}

// dp[i][j] = dp[i-1][j]+dp[i][j-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i-1][j] == 0 {
				dp[i][j] += dp[i-1][j]
			}
			if obstacleGrid[i][j-1] == 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
