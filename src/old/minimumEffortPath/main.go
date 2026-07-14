package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	res := minimumEffortPath([][]int{[]int{1, 2, 2}, []int{3, 8, 2}, []int{5, 3, 5}})
	fmt.Println(res)
}

func minimumEffortPath(heights [][]int) int {
	l := len(heights)
	arrived := map[string]int{}
	var dp func(x, y int) int
	dp = func(x, y int) int {
		if x == 0 && y == 0 {
			return 0
		}
		key := strconv.Itoa(x) + string(',') + strconv.Itoa(y)
		fmt.Println(key)
		_, ok := arrived[key]
		if ok {
			return arrived[key]
		}
		min, last, cur := math.MaxInt64, 0, 0
		if x-1 >= 0 {
			cur = int(math.Abs(float64(heights[x][y] - heights[x-1][y])))
			last = dp(x-1, y)
			fmt.Println(x, y)
			min = getMin(last, cur, min)
		}
		if x+1 < l {
			cur = int(math.Abs(float64(heights[x][y] - heights[x+1][y])))
			last = dp(x+1, y)
			fmt.Println(x, y)
			min = getMin(last, cur, min)
		}
		if y-1 >= 0 {
			cur = int(math.Abs(float64(heights[x][y] - heights[x][y-1])))
			last = dp(x, y-1)
			fmt.Println(x, y)
			min = getMin(last, cur, min)
		}
		if y+1 < l {
			cur = int(math.Abs(float64(heights[x][y] - heights[x][y+1])))
			last = dp(x, y+1)
			fmt.Println(x, y)
			min = getMin(last, cur, min)
		}
		arrived[key] = min
		return min
	}
	return dp(l-1, l-1)
}

func getMin(last, cur, min int) int {
	fmt.Println(last, cur, min)
	if last < cur {
		last = cur
	}
	if min > last {
		min = last
	}
	return min
}
