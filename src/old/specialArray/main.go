package main

import "fmt"

func main() {
	res := specialArray([]int{3, 6, 7, 7, 0})
	fmt.Println(res)
}

func specialArray(nums []int) int {
	count := 0
	for ; count < len(nums)+1; count++ {
		innerCount := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] >= count {
				innerCount++
			}
		}
		if innerCount == count {
			return count
		}
	}
	return -1
}
