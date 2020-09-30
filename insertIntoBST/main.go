package main

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, nil, nil}}
	insertIntoBST(&root, 5)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	finded := false
	var build func(root *TreeNode)
	build = func(root *TreeNode) {
		if root.Left != nil {
			build(root.Left)
		}
		if root.Val > val && !finded {
			newNode := &TreeNode{val, root.Left, nil}
			root.Left = newNode
			finded = true
			return
		}
		if root.Right != nil {
			build(root.Right)
		}
	}
	build(root)
	if !finded {
		var buildRight func(root *TreeNode)
		buildRight = func(root *TreeNode) {
			if root.Right == nil {
				newNode := &TreeNode{val, nil, nil}
				root.Right = newNode
				return
			}
			buildRight(root.Right)
		}
		buildRight(root)
	}
	return root
}
