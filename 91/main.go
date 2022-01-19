package main

import "fmt"

func main() {
	res := numDecodings("226")
	fmt.Println(res)
}

// dp[i] = dp[i-1] + dp[i-2]
func numDecodings(s string) int {
	l := len(s)
	dp := make([]int, l)
	if s[0] != '0' {
		dp[0] = 1
	} else {
		return 0
	}
	for i := 1; i < l; i++ {
		if s[i] != '0' {
			dp[i] += dp[i-1]
		}

		n := s[i-1] - '0'
		nn := s[i] - '0'
		if n == 1 || (n == 2 && nn < 7) {
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			} else {
				dp[i]++
			}
		}
	}
	return dp[l-1]
}
