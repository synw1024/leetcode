package main

import (
	"math"
)

var p1 = []int{31, 96, 73, 90, 15, 11, 1, 90, 72, 9, 30, 88}
var p2 = []int{87, 10, 3, 5, 76, 74, 38, 64, 16, 64, 93, 95, 60, 79, 54, 26, 30, 44, 64, 71}

func main() {
	assignTasks(p1, p2)
}

func assignTasks(servers []int, tasks []int) []int {
	serversLen := len(servers)
	tasksLen := len(tasks)
	status := make([]int, serversLen)
	res := []int{}
	for j, t := 0, 0; j < tasksLen; t++ {
		indexes := []int{}
		for i := 0; i < serversLen; i++ {
			if status[i] > 0 {
				status[i] -= 1
			}
			if status[i] > 0 {
				continue
			}
			indexes = append(indexes, i)
		}
		if len(indexes) == 0 {
			continue
		}
		for j <= t && j < tasksLen && len(indexes) > 0 {
			minWeight := math.MaxInt32
			minIndex := 0
			matchIndex := 0
			for i := 0; i < len(indexes); i++ {
				if servers[indexes[i]] < minWeight {
					minWeight = servers[indexes[i]]
					minIndex = indexes[i]
					matchIndex = i
				}
			}
			if matchIndex+1 < len(indexes) {
				indexes = append(indexes[:matchIndex], indexes[matchIndex+1:]...)
			} else {
				indexes = indexes[:matchIndex]
			}
			res = append(res, minIndex)
			status[minIndex] = tasks[j]
			j++
		}
	}
	return res
}
