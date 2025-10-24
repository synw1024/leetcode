package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := countAndSay(5)
	fmt.Println(res)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	ss := ""
	i, j := 0, 0
	for j < len(s) {
		if s[i] == s[j] {
			j++
		} else {
			ss += strconv.Itoa(j-i) + string(s[i])
			i = j
		}
	}
	ss += strconv.Itoa(j-i) + string(s[i])
	return ss
}
