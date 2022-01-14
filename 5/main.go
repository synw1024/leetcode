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

	start := 0
	length := 1
	for i := 0; i < l; i++ {
		j := i
		k := i
		for true {
			if j < 0 || k >= l || s[j] != s[k] {
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
			if j < 0 || k >= l || s[j] != s[k] {
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
