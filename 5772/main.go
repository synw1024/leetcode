package main

import "fmt"

func main() {
	res := isSumEqual("aaa", "a", "aaaa")
	fmt.Println(res)
}

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return transformStringToNumber(firstWord)+transformStringToNumber(secondWord) == transformStringToNumber(targetWord)
}

func transformStringToNumber(s string) rune {
	var n rune
	for i, v := range s {
		if i == 0 {
			n = v - rune('a')
			continue
		}
		n = n*10 + (v - rune('a'))
	}
	return n
}
