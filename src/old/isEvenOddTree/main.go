package main

import "fmt"

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := isEvenOddTree(&TreeNode{5, &TreeNode{9, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{1, &TreeNode{7, nil, nil}, nil}})
	fmt.Println(res)
}

func isEvenOddTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	var bfs func(nodes []*TreeNode, isEvenTier bool) bool
	bfs = func(nodes []*TreeNode, isEvenTier bool) bool {
		if len(nodes) == 0 {
			return true
		}
		if isEvenTier { // 偶数行
			if isEven(nodes[0].Val) {
				return false
			}
			for i := 1; i < len(nodes); i++ {
				if isEven(nodes[i].Val) || nodes[i].Val <= nodes[i-1].Val {
					return false
				}
			}
		} else { // 奇数行
			if !isEven(nodes[0].Val) {
				return false
			}
			for i := 1; i < len(nodes); i++ {
				if !isEven(nodes[i].Val) || nodes[i].Val >= nodes[i-1].Val {
					return false
				}
			}
		}

		temp := []*TreeNode{}
		for i := 0; i < len(nodes); i++ {
			if nodes[i].Left != nil {
				temp = append(temp, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				temp = append(temp, nodes[i].Right)
			}
		}
		return bfs(temp, !isEvenTier)
	}
	return bfs(queue, true)
}

func isEven(n int) bool {
	res := true
	if n%2 == 1 {
		return false
	}
	return res
}
