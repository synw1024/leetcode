package main

import "fmt"

func main() {
	res := firstUniqChar("leetcode")
	fmt.Println(res)
}

func firstUniqChar(s string) int {
	hashMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		v := s[i]
		if _, ok := hashMap[v]; ok {
			hashMap[v]++
		} else {
			hashMap[v] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if hashMap[v] == 1 {
			return i
		}
	}
	return -1
}
