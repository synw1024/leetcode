package main

import "fmt"

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree1 := TreeNode{0, &TreeNode{3, nil, nil}, &TreeNode{1, nil, &TreeNode{2, nil, nil}}}
	tree2 := TreeNode{0, &TreeNode{3, &TreeNode{2, nil, nil}, nil}, &TreeNode{1, nil, nil}}
	res := flipEquiv(&tree1, &tree2)
	fmt.Println(res)
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == root2 {
		return true
	} else if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) || flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left)
}
