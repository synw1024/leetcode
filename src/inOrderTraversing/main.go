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
	inOrderTraversing(&tree)
	fmt.Print("\n")
}

func inOrderTraversing(node *Node) {
	if node.Left == nil {
		fmt.Print(node.Val, " ")
		return
	}
	inOrderTraversing(node.Left)
	fmt.Print(node.Val, " ")
	if node.Right == nil {
		return
	}
	inOrderTraversing(node.Right)
}
