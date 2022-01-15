package main

import "fmt"

func main() {
	res := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
	res := []byte{}
	for i := 0; true; i++ {
		if i >= len(strs[0]) {
			return string(res)
		}
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return string(res)
			}
			if strs[0][i] != strs[j][i] {
				return string(res)
			}
		}
		res = append(res, strs[0][i])
	}
	return string(res)
}
