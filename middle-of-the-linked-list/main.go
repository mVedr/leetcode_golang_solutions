package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	s, f := head, head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}

func main() {

}
