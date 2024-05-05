package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	s, f := head, head.Next.Next
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	s.Next = s.Next.Next
	return head
}

func main() {

}
