package main

import "bytes"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return true
	}
	var buf bytes.Buffer
	a := head
	for a != nil {
		buf.WriteByte(byte(a.Val))
		a = a.Next
	}
	return ok(buf.String())
}

func ok(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {

}
