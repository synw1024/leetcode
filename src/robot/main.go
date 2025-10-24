package main

func main() {
	robot("URR", [][]int{[]int{2, 2}}, 3, 2)
}

func robot(command string, obstacles [][]int, x int, y int) bool {
	rx, ry := 0, 0
	validObstacles := [][]int{}
	for _, obstacle := range obstacles {
		if obstacle[0] <= x && obstacle[1] <= y {
			validObstacles = append(validObstacles, obstacle)
		}
	}
	for true {
		for _, v := range command {
			if v == 'U' {
				ry++
			} else {
				rx++
			}

			if rx == x && ry == y {
				return true
			}

			if rx > x || ry > y {
				return false
			}

			for _, obstacle := range validObstacles {
				if obstacle[0] == rx && obstacle[1] == ry {
					return false
				}
			}
		}
	}
	return false
}
