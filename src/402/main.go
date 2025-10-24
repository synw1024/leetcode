package main

import (
	"fmt"
	"strings"
)

func main() {
	num := removeKdigits("10", 2)
	fmt.Println(num)
}

func removeKdigits(num string, k int) string {
	if k == 0 {
		return num
	}

	for len(num) > 0 && k >= strings.Index(num, "0") && strings.Index(num, "0") != -1 {
		index := strings.Index(num, "0")
		num = num[index+1:]
		k -= index
		for len(num) > 0 && num[0] == '0' {
			num = num[1:]
		}
	}

	for k > 0 && len(num) > 0 {
		for i := 0; i < len(num); i++ {
			if i == len(num)-1 {
				num = num[:i]
				k--
				break
			}

			if num[i] > num[i+1] {
				num = num[:i] + num[i+1:]
				k--
				break
			}
		}
	}

	if num == "" {
		num = "0"
	}

	return num
}
