package main

import "fmt"

func main() {
	res := maxChunksToSorted([]int{1, 0, 2, 3, 4})
	fmt.Println(res)
}

func maxChunksToSorted(arr []int) int {
	count, max := 0, 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if max == i {
			count++
		}
	}
	return count
}
