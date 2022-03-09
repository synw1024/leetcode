package main

import "fmt"

func main() {
	res := findLongestChain([][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}})
	fmt.Println(res)
}

func findLongestChain(pairs [][]int) int {
	max, min, cache := pairs[0][0], pairs[0][1], map[int]int{}
	for _, v := range pairs {
		if max < v[0] {
			max = v[0]
		}
		if min > v[1] {
			min = v[1]
		}
	}

	var dp func(n int) int
	dp = func(left int) int {
		if left <= min {
			return 1
		}

		if v, ok := cache[left]; ok {
			return v
		}

		var max int
		for _, v := range pairs {
			if v[1] >= left {
				continue
			}
			res := dp(v[0]) + 1
			if res > max {
				max = res
			}
		}
		cache[left] = max
		return max
	}

	return dp(max)
}
