package main

import "fmt"

func main() {
	res := numDecodings("27")
	fmt.Println(res)
}

// dp[i] = dp[i-1] + dp[i-2]
func numDecodings(s string) int {
	l := len(s)
	dp := make([]int, l+1)
	if s[0] != '0' {
		dp[0] = 1
	}
	for i := 1; i <= l; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i-2 >= 0 {
			n := s[i-2] - '0'
			nn := s[i-1] - '0'
			if n == 1 || (n == 2 && nn < 7) {
				dp[i] += dp[i-2]
			}
		}

	}
	return dp[l]
}
