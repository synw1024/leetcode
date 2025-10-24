package main

import "fmt"

func main() {
	res := maxLengthBetweenEqualCharacters("cabbac")
	fmt.Println(res)
}

func maxLengthBetweenEqualCharacters(s string) int {
	max := -1
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && j-i-1 > max {
				max = j - i - 1
			}
		}
	}
	return max
}
