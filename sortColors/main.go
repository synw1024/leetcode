package main

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}

func sortColors(nums []int) {
	zeroArr, oneArr, twoArr := []int{}, []int{}, []int{}
	for _, v := range nums {
		if v == 0 {
			zeroArr = append(zeroArr, 0)
		} else if v == 1 {
			oneArr = append(oneArr, 1)
		} else {
			twoArr = append(twoArr, 2)
		}
	}
	zeroArr = append(zeroArr, oneArr...)
	zeroArr = append(zeroArr, twoArr...)
	nums = append(nums[:0], zeroArr...)
}
