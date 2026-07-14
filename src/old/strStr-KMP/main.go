package main

import "fmt"

func main() {
	res := strStr("mississippi", "pi")
	fmt.Println(res)
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	arr := kmpArr(needle)
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j == 0 {
				i++
			} else {
				i = i - arr[j-1] // (j-1) 已匹配的个数
				j = 0
			}
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

func kmpArr(s string) []int {
	res := []int{}
OUTER:
	for i := 0; i < len(s); i++ {
		ss := s[:i+1]
		for j := len(ss) - 1; j >= 0; j-- {
			head := ss[0:j]
			trail := ss[len(ss)-len(head) : len(ss)]
			if head == trail {
				res = append(res, len(head))
				continue OUTER
			}
		}
		res = append(res, 0)
	}
	return res
}
