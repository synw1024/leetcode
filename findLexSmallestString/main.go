package main

import "fmt"

func main() {
	res := findLexSmallestString()
	fmt.Println(res)
}

func findLexSmallestString(s string, a int, b int) string {
	if b%2 == 0 {
		ss := []int{}
		for i, v := range s {
			if i%2 == 0 {
				ss = append(ss, int(v))
				continue
			}
			ss = append(ss, int(v)%a)
		}
		if len(s)%b == 0 {
			ss
		}
	}
}

func findMinHead(nums []int, interval int) int {
	min, equalIndexs := nums[0], []int{0}
	for i := 0; i < len(nums); i += interval {
		if nums[i] < min {
			min = nums[i]
			equalIndexs = []int{i}
		} else if nums[i] == min {
			equalIndexs = append(equalIndexs, i)
		}
	}
	if len(equalIndexs) == 1 {
		return min
	}
	for ii := 1; ii < len(nums); ii++ {
		min, minIndexs := nums[equalIndexs[0]+ii], []int{equalIndexs[0] + ii}
		for _, v := range equalIndexs {
			if nums[v+ii] < min {
				min = nums[v+ii]
				minIndexs = []int{v + ii}
			}
		}
	}
	return 0
}
