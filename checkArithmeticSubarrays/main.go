package main

import (
	"fmt"
	"sort"
)

func main() {
	res := checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5})
	fmt.Println(res)
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := []bool{}
	for i, vl := range l {
		vr := r[i]
		temp := append([]int{}, nums[vl:vr+1]...)
		sort.Ints(temp)
		res = append(res, isTrue(temp))
	}
	return res
}

func isTrue(arr []int) bool {
	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}
	return true
}
