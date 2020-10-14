package main

import (
	"fmt"
	"strconv"
)

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree1 := TreeNode{1234, nil, nil}
	res := findMode(&tree1)
	fmt.Println(res)
}

func findMode(root *TreeNode) []int {
	nums := []int{}
	if root == nil {
		return nums
	}
	m := make(map[string]int)
	bfs(root, m)
	fmt.Println(m)
	max := 0
	for k, v := range m {
		n, err := strconv.Atoi(k)
		if err != nil {
			fmt.Println("error")
		}
		if v > max {
			max = v
			nums = []int{n}
		} else if v == max {
			nums = append(nums, n)
		}
	}
	return nums
}

func bfs(node *TreeNode, counter map[string]int) {
	s := strconv.Itoa(node.Val)
	_, ok := counter[s]
	if ok {
		counter[s]++
	} else {
		counter[s] = 1
	}
	if node.Left != nil {
		bfs(node.Left, counter)
	}
	if node.Right != nil {
		bfs(node.Right, counter)
	}
}
