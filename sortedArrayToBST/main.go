package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	res := &TreeNode{nums[len(nums)/2], nil, nil}
	res.Left = sortedArrayToBST(nums[0 : len(nums)/2])
	if len(nums)/2+1 < len(nums) {
		res.Right = sortedArrayToBST(nums[len(nums)/2+1:])
	}
	return res
}
