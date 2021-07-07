package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := longestPalindrome("cbbd")
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	l := len(s)

	dp := make([][3]string, l)
	dp[0][0] = "0,0"
	dp[0][1] = "0,1"
	dp[0][2] = "0,-1"

	for i := 1; i < l; i++ {
		dp[i][0] = strconv.Itoa(i) + "," + strconv.Itoa(i)
		dp[i][1] = max(s, strconv.Itoa(i)+","+strconv.Itoa(i+1), palindromeIndex(dp[i-1][1], i, s))
		dp[i][2] = max(s, dp[i-1][1], dp[i-1][2])
	}

	indexes := strings.Split(max(s, dp[l-1][1], dp[l-1][2]), ",")
	start := unsafeToNum(indexes[0])
	end := unsafeToNum(indexes[1])
	return s[start:end]
}

func palindromeIndex(indexesString string, index int, source string) string {
	indexes := strings.Split(indexesString, ",")
	start := unsafeToNum(indexes[0])

	ss := source[start : index+1]
	if !isPalindrome(ss) {
		return indexesString
	}
	return strconv.Itoa(start) + "," + strconv.Itoa(index+1)
}

func max(s, s1, s2 string) string {
	indexes1 := strings.Split(s1, ",")
	indexes2 := strings.Split(s2, ",")
	start1 := unsafeToNum(indexes1[0])
	start2 := unsafeToNum(indexes2[0])
	end1 := unsafeToNum(indexes1[1])
	end2 := unsafeToNum(indexes2[1])

	if end1-start1 > end2-start2 {
		return strconv.Itoa(start1) + "," + strconv.Itoa(end1)
	} else if end1-start1 == end2-start2 {
		if start1 < start2 {
			return strconv.Itoa(start1) + "," + strconv.Itoa(end1)
		}
		return strconv.Itoa(start2) + "," + strconv.Itoa(end2)
	}
	return strconv.Itoa(start2) + "," + strconv.Itoa(end2)
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

func unsafeToNum(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
