package main

import "fmt"

func main() {
	input := []int{2, 0, 2, 1, 1, 0}
	sortColors(input)
	fmt.Println(input)
}

func sortColors(nums []int) {
	l := len(nums)
	for i, start := 0, -1; i < l; i++ {
		if nums[i] == 0 && i != 0 {
			if start != -1 {
				temp := append([]int{}, nums[start:i]...)
				nums = append(nums[:start], nums[i:]...)
				nums = append(nums, temp...)
				i = start
				start = -1
			}
			nums = append(nums[:i], nums[i+1:]...)
			temp := append([]int{0}, nums...)
			nums = append(nums[:0], temp...)
		} else if nums[i] == 2 && start == -1 {
			start = i
		} else if nums[i] == 1 && start != -1 {
			temp := append([]int{}, nums[start:i]...)
			nums = append(nums[:start], nums[i:]...)
			nums = append(nums, temp...)
			i = start
			start = -1
		}
	}
}
