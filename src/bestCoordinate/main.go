package main

import (
	"fmt"
	"math"
)

func main() {
	res := bestCoordinate([][]int{[]int{28, 6, 30}, []int{23, 16, 0}, []int{21, 42, 22}, []int{50, 33, 34}, []int{14, 7, 50}, []int{40, 31, 4}, []int{39, 45, 17}, []int{46, 21, 12}, []int{45, 36, 45}, []int{35, 43, 43}, []int{29, 41, 48}, []int{22, 27, 5}, []int{42, 44, 45}, []int{10, 49, 50}, []int{47, 43, 26}, []int{40, 36, 25}, []int{10, 25, 6}, []int{27, 30, 30}, []int{50, 35, 20}, []int{11, 0, 44}, []int{34, 29, 28}}, 38)
	fmt.Println(res)
}

func bestCoordinate(towers [][]int, radius int) []int {
	var max float64
	var maxX, maxY int
	for y := 0; y < 51; y++ {
		for x := 0; x < 51; x++ {
			var sum float64
			for _, v := range towers {
				d := getDistance(x, y, v[0], v[1])
				if d > float64(radius) {
					continue
				}
				val := math.Floor(float64(v[2]) / (1 + d))
				sum += val
			}
			if (x == 45 && y == 36) || (x == 42 && y == 44) {
				fmt.Println(x, y, sum)
			}
			if max < sum {
				max = sum
				maxX = x
				maxY = y
			} else if max == sum {
				if x < maxX {
					max = sum
					maxX = x
					maxY = y
				} else if maxX == x && y < maxY {
					max = sum
					maxX = x
					maxY = y
				}
			}
		}
	}
	return []int{maxX, maxY}
}

func getDistance(x, y, tx, ty int) float64 {
	xx := float64(x)
	yy := float64(y)
	txx := float64(tx)
	tyy := float64(ty)
	return math.Sqrt(math.Pow(math.Abs(txx-xx), 2) + math.Pow(math.Abs(tyy-yy), 2))
}
