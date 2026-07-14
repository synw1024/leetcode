package main

import "fmt"

func main() {
	res := slowestKey([]int{12, 23, 36, 46, 62}, "spuda")
	fmt.Println(string(res))
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	max, key := 0, 'a'
	for i, v := range keysPressed {
		prev := 0
		if i != 0 {
			prev = releaseTimes[i-1]
		}
		diffTime := releaseTimes[i] - prev
		if max < diffTime {
			max = diffTime
			key = v
		} else if max == diffTime && v > key {
			max = diffTime
			key = v
		}
	}
	return byte(key)
}
