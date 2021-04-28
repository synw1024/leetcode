package main

import (
	"fmt"
	"math"
)

func main() {
	res := reverse(123)
	fmt.Println(res)
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		r := x % 10
		x /= 10
		if res > math.MaxInt32/10 || res == math.MaxInt32/10 && r > 7 {
			return 0
		}
		if res < -math.MaxInt32/10 || res == -math.MaxInt32/10 && r < -8 {
			return 0
		}
		res = res*10 + r
	}
	return res
}
