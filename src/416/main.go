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
		for j := 0; j <= half; j++ {
			if j == 0 {
				row = append(row, true)
				continue
			}
			row = append(row, false)
		}
		cache = append(cache, row)
	}
	cache[0][nums[0]] = true

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= half; j++ {
			if j >= nums[i] {
				cache[i][j] = cache[i-1][j] || cache[i-1][j-nums[i]]
			} else {
				cache[i][j] = cache[i-1][j]
			}
			if j == half && cache[i][j] {
				return true
			}
		}
	}
	return false
}
