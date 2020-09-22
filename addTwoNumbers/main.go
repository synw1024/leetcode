package main

import "fmt"

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{1, &ListNode{8, nil}}
	l2 := ListNode{0, nil}
	res := addTwoNumbers(&l1, &l2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{0, nil}
	point := &res
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			point.Val += l1.Val + l2.Val
		} else if l1 != nil {
			point.Val += l1.Val
		} else {
			point.Val += l2.Val
		}

		if point.Val >= 10 {
			point.Val = point.Val % 10
			point.Next = &ListNode{1, nil}
		} else if (l1 != nil && l1.Next != nil) || (l2 != nil && l2.Next != nil) {
			point.Next = &ListNode{0, nil}
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		point = point.Next
	}
	return &res
}
