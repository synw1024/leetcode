package main

import "fmt"

func main() {
	res := minimumOperations("rrryyyrryyyrr")
	fmt.Println(res)
}

func minimumOperations(leaves string) int {
	var solve func(index, state int) int
	var cache [][]int

	for i := 0; i < 3; i++ {
		for j := 0; j < len(leaves); j++ {
			cache[i][j] = -1
		}
	}

	solve = func(index, state int) int {
		if index == 0 {
			return isY(leaves[0])
		}

		if index < state {
			return int(^uint(0) >> 1)
		}

		if state == 2 {
			res1 := cache[1][index-1]
			if res1 == -1 {
				res1 = solve(index-1, 1)
			}

			res2 := cache[2][index-1]
			if res2 == -1 {
				res2 = solve(index-1, 2)
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

			if res2 < res1 {
				return res2 + isR(leaves[index])
			}
			return res1 + isR(leaves[index])
		} else {
			return solve(index-1, 0) + isY(leaves[index])
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
