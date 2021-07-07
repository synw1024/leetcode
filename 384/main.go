package main

import "math/rand"

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	temp := []int{}
	temp = append(temp, this.nums...)
	l := len(this.nums)

	for i := l - 1; i >= 0; i-- {
		index := rand.Int() % (i + 1)
		temp[i], temp[index] = temp[index], temp[i]
	}

	return temp
}
