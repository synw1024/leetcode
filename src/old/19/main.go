package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := []*ListNode{}
	for head != nil {
		temp = append(temp, head)
		head = head.Next
	}
	if len(temp) == 1 || len(temp) == 0 {
		return nil
	}
	if n == len(temp) {
		return temp[1]
	}
	temp[len(temp)-n-1].Next = temp[len(temp)-n].Next
	return temp[0]
}
