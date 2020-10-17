package main

import (
	"fmt"
)

func main() {
	res := canPartition([]int{1, 5, 11, 5})
	fmt.Println(res)
}

func canPartition(nums []int) bool {
	sum, max := 0, 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}
	if sum%2 != 0 {
		return false
	}
	half := sum / 2
	if max > half {
		return false
	}

	cache := [][]bool{}
	for i := 0; i < len(nums); i++ {
		row := []bool{}
		for j := 0; j < half+1; j++ {
			if j == 0 {
				row = append(row, true)
				continue
			}
			row = append(row, false)
		}
		cache = append(cache, row)
	}

	var dp func(ns []int, target int) bool
	dp = func(ns []int, target int) bool {
		if target == 0 {
			return true
		}
		if target < 0 {
			return false
		}
		for i, v := range ns {
			temp := append([]int{}, ns[:i]...)
			temp = append(temp, ns[i+1:]...)
			if dp(temp, target-v) {
				return true
			}
		}
		return false
	}
	return dp(nums, half)
}
