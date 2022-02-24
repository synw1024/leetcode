package main

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	faster := head
	if faster.Next == head {
		return true
	}
	for head != nil {
		if faster.Next == nil || faster.Next.Next == nil {
			return false
		}
		if head == faster.Next || head == faster.Next.Next {
			return true
		}
		head = head.Next
		faster = faster.Next.Next
	}
	return false
}
