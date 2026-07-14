package main

import "fmt"

func main() {
	res := mostPoints([][]int{
		[]int{1, 1},
		[]int{2, 2},
		[]int{3, 3},
		[]int{4, 4},
		[]int{5, 5},
	})
	fmt.Println(res)
}

func mostPoints(questions [][]int) int64 {
	l := len(questions)
	dp := make([]int64, l)
	dp[0] = int64(questions[0][0])

	var max int64 = dp[0]
	for i := 1; i < l; i++ {
		var max2 int64 = 0
		for j := 0; j < i; j++ {
			if questions[j][1]+j < i {
				if max2 < dp[j] {
					max2 = dp[j]
				}
			}
		}
		dp[i] = max2 + int64(questions[i][0])
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}
