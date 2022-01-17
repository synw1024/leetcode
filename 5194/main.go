package main

import (
	"fmt"
)

func main() {
	res := minMoves(10, 4)
	fmt.Println(res)
}

func minMoves(target int, maxDoubles int) int {
	count := 0
	for target > 1 {
		if target%2 == 0 && maxDoubles > 0 {
			target /= 2
			maxDoubles--
			count++
		} else if maxDoubles > 0 {
			target--
			count++
		} else {
			count += target - 1
			break
		}
	}
	return count
}
