package main

import (
	"fmt"
	"math"
)

func main() {
	res := divide(-2147483648, -1)
	fmt.Println(res)
}

func divide(dividend int, divisor int) int {
	count, isNagetive := 0, divisor < 0 && dividend > 0 || divisor > 0 && dividend < 0
	_dividend := math.Abs(float64(dividend))
	_divisor := math.Abs(float64(divisor))

	if int(_divisor) == 1 {
		if isNagetive {
			return -int(_dividend)
		}
		if int(_dividend) > 2147483647 {
			return 2147483647
		}
		return int(_dividend)
	}

	for true {
		if _dividend < _divisor {
			if isNagetive {
				return -count
			}
			return count
		}
		_dividend -= _divisor
		count++
	}
	return 0
}
