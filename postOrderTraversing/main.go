package main

import "fmt"

// Node struct
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func main() {
	tree := Node{1, &Node{2, &Node{4, nil, nil}, &Node{5, nil, nil}}, &Node{3, &Node{6, nil, nil}, &Node{7, nil, nil}}}
	postOrderTraversing(&tree)
	fmt.Print("\n")
}

func postOrderTraversing(node *Node) {
	if node.Left == nil && node.Right == nil {
		fmt.Print(node.Val, " ")
		return
	}
	if node.Left != nil {
		postOrderTraversing(node.Left)
	}
	if node.Right != nil {
		postOrderTraversing(node.Right)
	}
	fmt.Print(node.Val, " ")
}
