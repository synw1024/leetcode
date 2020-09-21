package main

import "fmt"

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := flipEquiv()
	fmt.Println(res)
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {

}
