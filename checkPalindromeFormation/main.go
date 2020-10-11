package main

import "fmt"

func main() {
	res := checkPalindromeFormation("abda", "acmc")
	fmt.Println(res)
}

func checkPalindromeFormation(a string, b string) bool {
	i := len(a)/2 - 1
	j := i + 2
	if len(a)%2 == 0 {
		i = len(a)/2 - 1
		j = i + 1
	}
	ti, tj := i, j

	resA := true
	for ti >= 0 {
		if a[ti] != a[tj] {
			if resA = reConcat(b[:ti+1]+a[ti+1:], ti, tj) || reConcat(a[:tj]+b[tj:], ti, tj); resA {
				return true
			}
			break
		}
		ti--
		tj++
	}

	ti, tj = i, j
	resB := true
	for ti >= 0 {
		if b[ti] != b[tj] {
			if resB = reConcat(a[:ti+1]+b[ti+1:], ti, tj) || reConcat(b[:tj]+a[tj:], ti, tj); resB {
				return true
			}
			break
		}
		ti--
		tj++
	}

	return resA || resB
}

func reConcat(s string, i, j int) bool {
	for i >= 0 {
		if s[i] != s[j] {
			return false
		}
		i--
		j++
	}
	return true
}
