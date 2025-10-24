package main

import (
	"fmt"
	"sort"
)

func main() {
	res := []int{1, 2, 3, 0, 0, 0}
	merge(res, 3, []int{2, 5, 6}, 3)
	fmt.Println(res)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}
