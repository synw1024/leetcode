package main

import (
	"fmt"
)

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		if index, ok := hashTable[target-v]; ok {
			return []int{index, i}
		}
		hashTable[v] = i
	}
	return nil
}
