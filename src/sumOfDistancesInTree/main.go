package main

import "fmt"

// TreeNode struct
type TreeNode struct {
	Val      int
	children []*TreeNode
}

func main() {
	res := sumOfDistancesInTree(6, [][]int{[]int{0, 1}, []int{0, 2}, []int{2, 3}, []int{2, 4}, []int{2, 5}})
	fmt.Println(res)
}

func sumOfDistancesInTree(N int, edges [][]int) []int {
	root := TreeNode{edges[0][0], nil}
}

func pickNodes() {

}
