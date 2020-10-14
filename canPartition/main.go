package main

import (
	"fmt"
)

func main() {
	res := canPartition([]int{1, 5, 11, 5})
	fmt.Println(res)
}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	half := sum / 2

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
