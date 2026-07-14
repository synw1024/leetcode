package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	res := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(res)
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; {
		for i < len(s) && !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
		}
		for j >= 0 && !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
		}
		if i > j {
			return true
		}
		si := strings.ToLower(string(s[i]))
		sj := strings.ToLower(string(s[j]))
		if si == sj {
			i++
			j--
		} else {
			fmt.Println(s[i], s[j])
			return false
		}
	}
	return true
}
