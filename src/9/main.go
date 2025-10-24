package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := isPalindrome(121)
	fmt.Println(res)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
