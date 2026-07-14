package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	res := trimMean([]int{4, 8, 4, 10, 0, 7, 1, 3, 7, 8, 8, 3, 4, 1, 6, 2, 1, 1, 8, 0, 9, 8, 0, 3, 9, 10, 3, 10, 1, 10, 7, 3, 2, 1, 4, 9, 10, 7, 6, 4, 0, 8, 5, 1, 2, 1, 6, 2, 5, 0, 7, 10, 9, 10, 3, 7, 10, 5, 8, 5, 7, 6, 7, 6, 10, 9, 5, 10, 5, 5, 7, 2, 10, 7, 7, 8, 2, 0, 1, 1})
	fmt.Println(res)
}

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	delLen := int(math.Floor(float64(len(arr))*0.05 + 0.5))
	sum := 0
	for i := delLen; i < len(arr)-delLen; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(len(arr)-2*delLen)
}
