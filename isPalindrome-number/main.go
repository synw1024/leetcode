package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func isPalindrome(head *ListNode) bool {
	faster := head
	slower := head
	var prev *ListNode
	endFound, res := false, true
	for faster != nil {
		if !endFound && faster.Next != nil && faster.Next.Next != nil {
			faster = faster.Next.Next
			next := slower.Next
			slower.Next = prev
			prev = slower
			slower = next
		} else if !endFound && faster.Next == nil {
			endFound = true
			faster = slower.Next
			if prev == nil {
				break
			}
			next := prev.Next
			prev.Next = slower
			slower = prev
			prev = next
		} else if !endFound && faster.Next.Next == nil {
			endFound = true
			faster = slower.Next
		} else {
			if slower.Val != faster.Val {
				res = false
			}
			faster = faster.Next
			if prev == nil {
				break
			}
			next := prev.Next
			prev.Next = slower
			slower = prev
			prev = next
		}
		fmt.Println(slower.Val, faster.Val)
	}
	return res
}
