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
	res := 0
	if divisor == 0 {
		res = math.MaxInt32
	}
	if divisor == 1 {
		res = dividend
	}
	if divisor == -1 {
		res = -dividend
	}
	isNagative := false
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		isNagative = true
	}

	res = div(abs(dividend), abs(divisor))

	if isNagative {
		return -res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
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

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
