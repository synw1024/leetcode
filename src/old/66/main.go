package main

import "fmt"

func main() {
	res := plusOne([]int{1, 2, 3})
	fmt.Println(res)
}

func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] > 9 {
			digits[i] -= 10
			if i == 0 {
				digits = append([]int{1}, digits...)
			} else {
				digits[i-1]++
			}
		}
	}
	return digits
}
