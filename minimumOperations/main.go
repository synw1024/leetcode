package main

import "fmt"

func main() {
	res := minimumOperations("rrryyyrryyyrr")
	fmt.Println(res)
}

func minimumOperations(leaves string) int {
	countArr := []int{0}
	for _, v := range leaves {
		if v == 'r' {
			if countArr[len(countArr)-1] > 0 {
				countArr = append(countArr, -1)
			} else {
				countArr[len(countArr)-1]--
			}
		} else {
			if countArr[len(countArr)-1] < 0 {
				countArr = append(countArr, 1)
			} else {
				countArr[len(countArr)-1]++
			}
		}
	}

	count := 0
	if countArr[0] > 0 {
		count++
		if countArr[0] > 1 {
			countArr[0]--
		} else {
			countArr = countArr[1:]
			if len(countArr) == 1 {
				count++
			}
		}
	}
	if countArr[len(countArr)-1] > 0 {
		count++
		if countArr[len(countArr)-1] > 1 {
			countArr[len(countArr)-1]--
		} else {
			countArr = countArr[:len(countArr)-1]
			if len(countArr) == 1 {
				count++
			}
		}
	}

	max, maxIndex := countArr[0], 0
	for i := 1; i < len(countArr); i++ {
		if max < countArr[i] {
			max = countArr[i]
			maxIndex = i
		}
	}
}
