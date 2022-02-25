package main

import "fmt"

func main() {
	res := containsDuplicate([]int{1, 2, 3})
	fmt.Println(res)
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v] = true
	}
	return false
}
