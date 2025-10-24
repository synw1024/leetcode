package main

import (
	"fmt"
)

func main() {
	res := divideString("abcdefghij", 3, 'x')
	fmt.Println(res)
}

func divideString(s string, k int, fill byte) []string {
	res := []string{}
	for i := range s {
		if i%k == 0 {
			res = append(res, s[i:i+1])
		} else {
			res[len(res)-1] += s[i : i+1]
		}
	}
	diff := k - len(res[len(res)-1])
	if diff > 0 {
		for i := 0; i < diff; i++ {
			res[len(res)-1] += string([]byte{fill})
		}
	}
	return res
}
