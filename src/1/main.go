package main

import (
	"fmt"
)

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		idx, ok := m[v]
		fmt.Println(idx, ok)
		if ok {
			return []int{idx, i}
		} else {
			m[target-v] = i
		}
	}
	return []int{}
}
