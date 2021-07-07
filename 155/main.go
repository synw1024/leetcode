package main

type MinStack struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.nums = append(this.nums, val)
}

func (this *MinStack) Pop() {
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	min := this.nums[0]
	l := len(this.nums)
	for i := 1; i < l; i++ {
		if min > this.nums[i] {
			min = this.nums[i]
		}
	}
	return min
}
