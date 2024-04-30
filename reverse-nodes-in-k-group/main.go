package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var prev, curr, next *ListNode
	ll := 0
	ptr := head
	for ptr != nil {
		ptr = ptr.Next
		ll++
	}
	if ll < k {
		return head
	}
	curr = head
	next = curr.Next
	ll = 0
	for ll < k {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		ll++
	}
	head.Next = reverseKGroup(curr, k)
	return prev
}

func main() {

}
