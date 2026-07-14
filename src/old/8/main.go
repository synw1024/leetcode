package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
	res := myAtoi("42")
	fmt.Println(res)
}

func myAtoi(s string) int {
	b := []byte(s)
	b2 := []byte{}
	isSpace := true
	sign := 0
	meetSign := false
	for _, v := range b {
		if isSpace && v == ' ' {
			continue
		} else {
			isSpace = false
		}
		if !meetSign && (v == '+' || v == '-') {
			if v == '+' {
				sign = 1
			} else {
				sign = -1
			}
			meetSign = true
			continue
		}
		meetSign = true
		if unicode.IsDigit(rune(v)) {
			b2 = append(b2, v)
		} else {
			break
		}
	}
	if len(b2) == 0 {
		return 0
	}
	n := 0
	for _, v := range b2 {
		vv := int(v - '0')
		if sign != -1 && (n > math.MaxInt32/10 || (n == math.MaxInt32/10 && vv > 7)) {
			return math.MaxInt32
		}
		if sign == -1 && (-n < -math.MaxInt32/10 || (-n == -math.MaxInt32/10 && -vv < -8)) {
			return -math.MaxInt32 - 1
		}
		n = n*10 + vv
	}
	if sign == -1 {
		return -n
	}
	return n
}
