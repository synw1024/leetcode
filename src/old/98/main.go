package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	node := root
	prev := -math.MaxInt64
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val <= prev {
			return false
		}
		prev = node.Val
		node = node.Right
	}
	return true
}
