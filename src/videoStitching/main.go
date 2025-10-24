package main

import (
	"fmt"
)

func main() {
	// res := videoStitching([][]int{[]int{0, 2}, []int{4, 6}, []int{8, 10}, []int{1, 9}, []int{1, 5}, []int{5, 9}}, 10)
	res := videoStitching([][]int{[]int{0, 1}, []int{1, 2}}, 5)
	fmt.Println(res)
}

func videoStitching(clips [][]int, T int) int {
	var dp func(clips [][]int, T int) int
	dp = func(clips [][]int, T int) int {
		if T == 0 {
			return 0
		}
		min := 101
		for i, v := range clips {
			if v[1] >= T {
				temp := append([][]int{}, clips[:i]...)
				temp = append(temp, clips[i+1:]...)
				res := dp(temp, v[0])
				if res == -1 {
					continue
				}
				if res+1 < min {
					min = res + 1
				}
			}
		}
		if min == 101 {
			return -1
		}
		return min
	}
	return dp(clips, T)
}
