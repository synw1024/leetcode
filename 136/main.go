package main

import "fmt"

func main() {
	res := singleNumber([]int{1, 3, 1, -1, 3})
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	var arr []int
	for _, v := range nums {
		arr = contains(arr, v)
	}
	return arr[0]
}

func contains(arr []int, n int) []int {
	if len(arr) == 0 {
		return []int{n}
	}
	var res []int
	for i, v := range arr {
		if v == n {
			if i == len(arr)-1 {
				res = arr[:i]
			} else {
				res = append(arr[:i], arr[i+1:]...)
			}
			return res
		}
	}
	return append(arr, n)
}
