package main

import (
	"fmt"
	"strings"
)

func main() {
	res := partitionLabels("ababcbacadefegdehijhklij")
	fmt.Println(res)
}

func partitionLabels(S string) []int {
	count, k := []int{}, 0
	j := 0
	for i := 0; i < len(S); i++ {
		_j := strings.LastIndex(S, string(S[i]))
		if _j > j {
			j = _j
		}
		if i == j {
			count = append(count, j-k+1)
			k = i + 1
		}
	}
	return count
}
