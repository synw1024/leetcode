package main

import (
	"fmt"
)

func main() {
	res := maxSatisfied([]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3)
	fmt.Println(res)
}

func maxSatisfied(customers []int, grumpy []int, X int) int {
	var product []int
	for i, v := range customers {
		product = append(product, v*grumpy[i])
	}

	max := 0
	maxIndex := 0
	for i := 0; i < len(product)-X+1; i++ {
		sum := 0
		for ii := 0; ii < X; ii++ {
			sum += product[i+ii]
		}

		if max < sum {
			max = sum
			maxIndex = i
		}
	}

	for i := 0; i < X; i++ {
		grumpy[maxIndex+i] = 0
	}

	sum := 0
	for i, v := range customers {
		if grumpy[i] == 0 {
			sum += v
		}
	}

	return sum
}
