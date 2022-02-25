package main

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// num := sumOfLeftLeaves()
	// fmt.Println(num)
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0)
}

func dfs(node *TreeNode, sum int) int {
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		sum += node.Left.Val
	}

	if node.Left != nil {
		sum = dfs(node.Left, sum)
	}

	if node.Right != nil {
		sum = dfs(node.Right, sum)
	}

	return sum
}
