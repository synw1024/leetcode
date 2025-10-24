package main

import "fmt"

func main() {
	res := minOperations([]string{"d1/", "../", "../", "../"})
	fmt.Println(res)
}

func minOperations(logs []string) int {
	n := 0
	for _, v := range logs {
		if v == "../" {
			if n > 0 {
				n--
			}
			continue
		} else if v == "./" {
			continue
		}
		n++
	}
	return n
}
