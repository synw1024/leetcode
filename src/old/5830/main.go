package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isThree(n int) bool {
	res := math.Sqrt(float64(n))
	numstr := fmt.Sprint(res)
	temp := strings.Split(numstr, ".")
	if len(temp) <= 1 {
		num, err := strconv.Atoi(temp[0])
		if err != nil {
			return false
		}
		if isPrime(num) {
			return true
		}
		return false
	}
	return false
}

func isPrime(num int) bool {
	if num <= 3 {
		return num > 1
	}
	// 不在6的倍数两侧的一定不是质数
	if num%6 != 1 && num%6 != 5 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 5; i <= sqrt; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
