package main

import (
	"fmt"
)

func main() {
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(res)
}

func maxArea(height []int) int {
	res := 0
	for i, j := 0, len(height)-1; i != j; {
		var isIMin bool
		var h int
		if height[i] < height[j] {
			h = height[i]
			isIMin = true
		} else {
			h = height[j]
		}
		area := (j - i) * h
		if area > res {
			res = area
		}
		if isIMin {
			i++
		} else {
			j--
		}
	}
	return res
}
