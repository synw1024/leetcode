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
	bfs([]*Node{&tree})
}

func bfs(queue []*Node) {
	first := queue[0]
	fmt.Print(first.Val, " ")
	fmt.Print("\n")
	if first.Left != nil {
		queue = append(queue, first.Left)
	}
	if first.Right != nil {
		queue = append(queue, first.Right)
	}
	queue = queue[1:]
	if len(queue) == 0 {
		return
	}
	bfs(queue)
}
