package main

import "fmt"

func main() {
	res := maxSubArray([]int{5, 4, -1, 7, 8})
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	max := nums[0]
	for i := 0; i < l; {
		for j := i + 1; j < l; {
			temp := sum(nums[i : j+1])
			if max < temp {
				max = temp
				j++
				continue
			}
			i = j
			break
		}
		i++
	}

	return max
}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
