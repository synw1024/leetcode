package main

import "fmt"

func main() {
	res := maximalNetworkRank(8, [][]int{[]int{0, 1}, []int{1, 2}, []int{2, 3}, []int{2, 4}, []int{5, 6}, []int{5, 7}})
	fmt.Println(res)
}

func maximalNetworkRank(n int, roads [][]int) int {
	hash := map[int]int{}
	for _, v := range roads {
		for _, vv := range v {
			if _, ok := hash[vv]; ok {
				hash[vv]++
			} else {
				hash[vv] = 1
			}
		}
	}

	max, max2, key := 0, 0, 0
	for k, v := range hash {
		if v > max {
			max = v
			key = k
		}
	}

	for k, v := range hash {
		if k == key {
			continue
		}
		if isConnect(roads, key, k) && v-1 > max2 {
			max2 = v - 1
		} else if !isConnect(roads, key, k) && v > max2 {
			max2 = v
		}
	}

	return max + max2
}

func isConnect(roads [][]int, key, key2 int) bool {
	for _, v := range roads {
		if (v[0] == key && v[1] == key2) || (v[0] == key2 && v[1] == key) {
			return true
		}
	}
	return false
}
