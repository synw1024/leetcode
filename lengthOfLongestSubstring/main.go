package main

import "fmt"

func main() {
	res := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	max, count, hash := 0, 0, map[byte]int{}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if ii, ok := hash[v]; ok {
			hash = map[byte]int{}
			if max < count {
				max = count
			}
			count = 0
			i = ii
		} else {
			hash[v] = i
			count++
		}
	}
	if count > max {
		return count
	}
	return max
}
