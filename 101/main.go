package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	res := isSymmetric(&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}})
	fmt.Println(res)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	a := []*TreeNode{root}
	for len(a) > 0 {
		fmt.Println(a)
		for i, j := 0, len(a)-1; i < j; {
			if a[i] == nil && a[j] == nil {
				i++
				j--
				continue
			}
			if a[i] == nil || a[j] == nil {
				return false
			}
			if a[i].Val != a[j].Val {
				return false
			}
			i++
			j--
		}
		l := len(a)
		for i := 0; i < l; i++ {
			if a[i] == nil {
				continue
			}
			a = append(a, a[i].Left, a[i].Right)
		}
		a = a[l:]
		fmt.Println(len(a))
	}
	return true
}
