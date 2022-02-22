package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	temp := []*TreeNode{root}
	res := [][]int{}
	for len(temp) > 0 {
		row := []int{}
		l := len(temp)
		for i := 0; i < l; i++ {
			if temp[i] == nil {
				continue
			}
			row = append(row, temp[i].Val)
			temp = append(temp, temp[i].Left, temp[i].Right)
		}
		if len(row) > 0 {
			res = append(res, row)
		}
		temp = temp[l:]
	}
	return res
}
