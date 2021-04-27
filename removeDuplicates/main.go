package main

import "fmt"

func main() {
	l := removeDuplicates([]int{1, 2, 2, 3, 3, 4, 5})
	fmt.Println(l)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	i := 0
	for j := 1; j < len(nums); {
		if nums[i] == nums[j] {
			j++
		} else {
			nums = append(nums[:i+1], nums[j:]...)
			i++
			j = i + 1
		}
	}
	nums = nums[:i+1]
	return len(nums)
}
