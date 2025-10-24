package main

import "fmt"

func main() {
	res := minOperationsMaxProfit([]int{10, 10, 6, 4, 7}, 3, 8)
	fmt.Println(res)
}

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	costs := []int{}
	runingCustomers := 0
	for i := 0; true; {
		if customers[0] > 4 {
			runingCustomers += 4
			costs = append(costs, runingCustomers*boardingCost-(i+1)*runningCost)
			if 1 == len(customers) {
				customers = append(customers, customers[0]-4)
			} else {
				customers[1] += customers[0] - 4
			}
		} else {
			runingCustomers += customers[0]
			costs = append(costs, (runingCustomers*boardingCost)-(i+1)*runningCost)
		}
		customers = customers[1:]
		if len(customers) == 0 {
			break
		}
		i++
	}

	max := costs[0]
	maxIndex := 0
	for i, v := range costs {
		if v > max {
			max = v
			maxIndex = i
		}
	}
	if max > 0 {
		return maxIndex + 1
	}
	return -1
}
