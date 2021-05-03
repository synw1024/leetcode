package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	n1, n2 := l1, l2
	for l1 != nil || l2 != nil {
		if l1.Next == nil {
			l1.Next = l2
			break
		}
		if l1.Next.Val <= l2.Val {
			l1 = l1.Next
		} else {
			n2 = l2
			l2 = l2.Next
			n2.Next = l1.Next
			l1.Next = n2
			if l2 == nil {
				break
			}
		}
	}
	return n1
}
