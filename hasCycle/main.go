package main

// ListNode struct
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle(head *ListNode) bool {
	pointerArr, pointer := []*ListNode{}, head
	for pointer != nil {
		for _, v := range pointerArr {
			if v == pointer {
				return true
			}
		}
		pointerArr = append(pointerArr, pointer)
		pointer = pointer.Next
	}
	return false
}
