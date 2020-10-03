package main

import (
	"fmt"
	"strings"
)

func main() {
	res := getHappyString(10, 100)
	fmt.Println(res)
}

func getHappyString(n int, k int) string {
	arr := []int{}
	for j := 0; j < n; j++ {
		if j%2 == 0 {
			arr = append(arr, 0)
		} else {
			arr = append(arr, 1)
		}
	}

	temp := append([]int{}, arr...)
	for i := 1; i < k; {
		if !addOne(&temp) {
			return ""
		}
		if !isDuplication(&temp) {
			arr = append([]int{}, temp...)
			i++
		}
	}

	res := []string{}
	for _, v := range arr {
		if v == 0 {
			res = append(res, "a")
		} else if v == 1 {
			res = append(res, "b")
		} else {
			res = append(res, "c")
		}
	}

	return strings.Join(res, "")
}

func addOne(_arr *[]int) bool {
	arr := *_arr
	arr[len(arr)-1]++
	for i := len(arr) - 1; i >= 0; i-- {
		if i == 0 && arr[i] > 2 {
			return false
		} else if arr[i] > 2 {
			arr[i] = 0
			arr[i-1]++
		}
	}
	return true
}

func isDuplication(_arr *[]int) bool {
	arr := *_arr
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return true
		}
	}
	return false
}
