package main

import "fmt"

func main() {
	res := maxDepth("1")
	fmt.Println(res)
}

func maxDepth(s string) int {
	max, count := 0, 0
	for _, v := range s {
		if v == '(' {
			count++
			if count > max {
				max = count
			}
		} else if v == ')' {
			count--
		}
	}
	return max
}
