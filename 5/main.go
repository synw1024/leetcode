package main

import (
	"fmt"
)

func main() {
	res := longestPalindrome("babad")
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	l := len(s)

	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	start := 0
	length := 1
	for i := 0; i < l; i++ {
		j := i
		k := i
		for true {
			if j < 0 || k >= l || !isPalindrome(s[j:k+1]) {
				break
			}
			if k-j+1 > length {
				start = j
				length = k - j + 1
			}
			j--
			k++
		}
		j = i
		k = i + 1
		for true {
			if j < 0 || k >= l || !isPalindrome(s[j:k+1]) {
				break
			}
			if k-j+1 > length {
				start = j
				length = k - j + 1
			}
			j--
			k++
		}
	}

	return s[start : start+length]
}

func isPalindrome(s string) bool {
	l := len(s)

	for i, j := 0, l-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
