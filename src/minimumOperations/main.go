package main

import (
	"fmt"
	"math"
)

func main() {
	res := minimumOperations("rrryyyrryyyrr")
	fmt.Println(res)
}

func minimumOperations(leaves string) int {
	var solve func(index, state int) int
	var cache [][]int

	for i := 0; i < 3; i++ {
		row := []int{}
		for j := 0; j < len(leaves); j++ {
			row = append(row, -1)
		}
		cache = append(cache, row)
	}

	solve = func(index, state int) int {
		if index < state {
			return math.MaxInt64
		}

		if index == 0 {
			return isY(leaves[0])
		}

		if state == 2 {
			res1 := cache[1][index-1]
			if res1 == -1 {
				res1 = solve(index-1, 1)
				cache[1][index-1] = res1
			}

			res2 := cache[2][index-1]
			if res2 == -1 {
				res2 = solve(index-1, 2)
				cache[2][index-1] = res2
			}

			if res2 < res1 {
				cache[2][index] = res2 + isY(leaves[index])
			} else {
				cache[2][index] = res1 + isY(leaves[index])
			}
			return cache[2][index]
		} else if state == 1 {
			res1 := cache[0][index-1]
			if res1 == -1 {
				res1 = solve(index-1, 0)
			}

			res2 := cache[1][index-1]
			if res2 == -1 {
				res2 = solve(index-1, 1)
			}

			if res2 < res1 {
				cache[1][index] = res2 + isR(leaves[index])
			} else {
				cache[1][index] = res1 + isR(leaves[index])
			}
			return cache[1][index]
		} else {
			res := cache[0][index-1]
			if res == -1 {
				res = solve(index-1, 0)
			}

			cache[0][index] = res + isY(leaves[index])
			return cache[0][index]
		}
	}
	return solve(len(leaves)-1, 2)
}

func isY(leave byte) int {
	if leave == 'y' {
		return 1
	}
	return 0
}

func isR(leave byte) int {
	if leave == 'r' {
		return 1
	}
	return 0
}
