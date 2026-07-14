package main

import "fmt"

func main() {
	res := reverseParentheses("(abcd)")
	fmt.Println(res)
}

func reverseParentheses(s string) string {
	bracketsIndex := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			bracketsIndex = append(bracketsIndex, i)
		} else if s[i] == ')' {
			index := bracketsIndex[len(bracketsIndex)-1]
			bracketsIndex = bracketsIndex[:len(bracketsIndex)-1]
			s = s[:index] + reverse(s[index+1:i]) + s[i+1:]
			i -= 2
		}
	}
	return s
}

func reverse(s string) string {
	ss := []rune(s)
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	}
	return string(ss)
}
