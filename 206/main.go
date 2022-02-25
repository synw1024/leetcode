package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		if next == nil {
			return head
		}
		head = next
	}
	return head
}
