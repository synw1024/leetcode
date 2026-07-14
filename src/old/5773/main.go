package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := maxValue("-13", 2)
	fmt.Println(res)
}

func maxValue(n string, x int) string {
	l := len(n)
	if n[0] != '-' {
		for i := 0; i < l; i++ {
			if n[i]-'0' < byte(x) {
				return n[:i] + strconv.Itoa(x) + n[i:]
			}
		}
	} else {
		for i := 1; i < l; i++ {
			if n[i]-'0' > byte(x) {
				return n[:i] + strconv.Itoa(x) + n[i:]
			}
		}
	}
	return n + strconv.Itoa(x)
}
