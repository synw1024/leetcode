package main

import "fmt"

func main() {
	res := strStr("mississippi", "issip")
	fmt.Println(res)
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		h := haystack[i]
		n := needle[j]
		if h == n {
			i++
			j++
		} else {
			i -= j - 1
			j = 0
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}
