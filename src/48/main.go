package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	i := len(nums) - k
	reverse(nums[:i])
	reverse(nums[i:])
	reverse(nums)
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
