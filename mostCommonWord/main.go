package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	res := mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"})
	fmt.Println(res)
}

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)

	re := regexp.MustCompile(`\W`)
	paragraph = re.ReplaceAllString(paragraph, "-")

	paragraphArr := strings.Split(paragraph, "-")

	m := make(map[string]int)
	for i := 0; i < len(paragraphArr); i++ {
		if paragraphArr[i] == "" {
			continue
		}

		_, ok := m[paragraphArr[i]]
		if !ok {
			m[paragraphArr[i]] = 1
		} else {
			m[paragraphArr[i]]++
		}
	}

	maxKey := ""
	for k, v := range m {
		if (maxKey == "" || m[maxKey] < v) && !contains(banned, k) {
			maxKey = k
		}
	}

	return maxKey
}

func contains(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
