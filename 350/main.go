package main

import (
	"fmt"
	"sort"
)

func main() {
	res := intersect([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Println(res)
}

func intersect(nums1 []int, nums2 []int) []int {
	var short, long []int
	if len(nums1) > len(nums2) {
		long = nums1
		short = nums2
	} else {
		long = nums2
		short = nums1
	}
	sort.Ints(long)
	sort.Ints(short)
	var res []int
	for l, s := 0, 0; l < len(long) && s < len(short); {
		if long[l] == short[s] {
			res = append(res, short[s])
			l++
			s++
		} else {
			if long[l] > short[s] {
				s++
			} else {
				l++
			}
		}
	}
	return res
}
