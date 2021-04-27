package main

import "fmt"

func main() {
	res := []int{0, 1, 0, 3, 12}
	moveZeroes(res)
	fmt.Println(res)
}

func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); {
		if nums[j] == 0 {
			j++
		} else if i != j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			i++
			j++
		}
	}
}
