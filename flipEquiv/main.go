package main

import "fmt"

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree1 := TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{7, nil, nil}, &TreeNode{8, nil, nil}}}, &TreeNode{3, &TreeNode{6, nil, nil}, nil}}
	tree2 := TreeNode{1, &TreeNode{3, nil, &TreeNode{6, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, &TreeNode{8, nil, nil}, &TreeNode{7, nil, nil}}}}
	res := flipEquiv(&tree1, &tree2)
	fmt.Println(res)
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	childrenEqual(root1)
}

func childrenEqual(node1 *TreeNode, node2 *TreeNode) bool {
	var arr1, arr2 []int
	if node1.Left != nil && node1.Right != nil && node2.Left != nil && node2.Right != nil {
		if node1.Left.Val == node2.Left.Val && node1.Right.Val == node2.Right.Val {
			return true
		} else if node1.Left.Val == node2.Right.Val && node1.Right.Val == node2.Left.Val {
			return true
		} else {
			return false
		}
	} else if (node1.Left != nil && node1.Right == nil || node1.Right != nil && node1.Left == nil) && (node2.Left != nil && node2.Right == nil || node2.Right != nil && node2.Left == nil) {
		var val1, val2 int
		if node1.Left != nil {
			val1 = node1.Left.Val
		} else {
			val1 = node1.Right.Val
		}
		if node2.Left != nil {
			val2 = node2.Left.Val
		} else {
			val2 = node2.Right.Val
		}
		return val1 == val2
	} else if node1.Left == nil && node1.Right == nil && node2.Left == nil && node2.Right == nil {
		return true
	}
	return false
}
