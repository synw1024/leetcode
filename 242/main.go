package main

import (
	"fmt"
	"sort"
)

func main() {
	res := isAnagram("r", "cat")
	fmt.Println(res)
	s := "哈合"
	fmt.Println([]rune(s))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var a1, a2 []int
	for _, v := range s {
		a1 = append(a1, int(v))
	}
	sort.Ints(a1)
	for _, v := range t {
		a2 = append(a2, int(v))
	}
	sort.Ints(a2)
	for i, v := range a1 {
		if v != a2[i] {
			return false
		}
	}
	return true
}
