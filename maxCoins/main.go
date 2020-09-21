package main

import "fmt"

func main() {
	res := maxCoins([]int{3, 1, 5, 8})
	fmt.Println(res)
}

func maxCoins(nums []int) int {
	for true {
		index := indexOf(&nums, 0)
		if index == -1 {
			break
		}
		nums = append(nums[:index], nums[index+1:]...)
	}
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)

	records := make([][]int, len(nums)+2)
	for i := 0; i < len(nums)+2; i++ {
		records[i] = make([]int, len(nums)+2)
		for j := 0; j < len(nums)+2; j++ {
			records[i][j] = -1
		}
	}

	res := solve(&nums, &records, 0, len(nums)-1)
	return res
}

func solve(nums *[]int, records *[][]int, l int, r int) int {
	if l >= r-1 {
		return 0
	}
	if (*records)[l][r] != -1 {
		return (*records)[l][r]
	}
	for i := l + 1; i < r; i++ {
		sum := solve(nums, records, l, i) + (*nums)[l]*(*nums)[i]*(*nums)[r] + solve(nums, records, i, r)
		if (*records)[l][r] < sum {
			(*records)[l][r] = sum
		}
	}
	return (*records)[l][r]
}

func indexOf(nums *[]int, n int) int {
	for i, v := range *nums {
		if v == n {
			return i
		}
	}
	return -1
}
