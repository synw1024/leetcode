package main

import "fmt"

// TreeNode struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := verticalTraversal(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}})
	fmt.Println(res)
}

func verticalTraversal(root *TreeNode) [][]int {
	posHash := map[int]map[int][]int{}
	var dfs func(node *TreeNode, posX, posY int)
	dfs = func(node *TreeNode, posX, posY int) {
		yHash, xOk := posHash[posX]
		if xOk {
			yArr, yOk := yHash[posY]
			if yOk {
				index := insertIndex(yArr, node.Val)
				temp := append([]int{node.Val}, yArr[index:]...)
				yHash[posY] = append(yArr[:index], temp...)
			} else {
				yHash[posY] = []int{node.Val}
			}
		} else {
			posHash[posX] = map[int][]int{}
			posHash[posX][posY] = []int{node.Val}
		}
		if node.Left != nil {
			dfs(node.Left, posX-1, posY+1)
		}
		if node.Right != nil {
			dfs(node.Right, posX+1, posY+1)
		}
	}
	dfs(root, 0, 0)

	keys1 := []int{}
	for k := range posHash {
		index := insertIndex(keys1, k)
		temp := append([]int{k}, keys1[index:]...)
		keys1 = append(keys1[:index], temp...)
	}
	res := [][]int{}
	for k := range keys1 {
		keys2 := []int{}
		for k2 := range posHash[keys1[k]] {
			index := insertIndex(keys2, k2)
			temp2 := append([]int{k2}, keys2[index:]...)
			keys2 = append(keys2[:index], temp2...)
		}

		temp := []int{}
		for kk := range keys2 {
			temp = append(temp[0:], posHash[keys1[k]][keys2[kk]]...)
		}
		res = append(res, temp)
	}
	return res
}

func insertIndex(arr []int, n int) int {
	for i, v := range arr {
		if n < v {
			return i
		}
	}
	return len(arr)
}
