package main

import (
	"fmt"
	"math"
)

func main() {
	res := divide(10, 0)
	fmt.Println(res)
}

func divide(dividend, divisor int) int {
	if divisor == 0 {
		return math.MaxInt32
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend
	}
	return div(dividend, divisor)
}

func div(dividend, divisor int) int {
	if dividend < divisor {
		return 0
	}

	tempDivisor, count := divisor, 1

	for tempDivisor+tempDivisor <= dividend {
		tempDivisor += tempDivisor
		count += count
	}

	return count + div(dividend-tempDivisor, divisor)
}
