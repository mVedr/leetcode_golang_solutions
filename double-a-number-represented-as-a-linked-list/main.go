package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ok(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func doubleIt(head *ListNode) *ListNode {
	newhead := ok(head)

	ptr := newhead
	carry := 0
	for ptr != nil {
		doubled := ptr.Val*2 + carry
		ptr.Val = doubled % 10
		carry = doubled / 10
		if ptr.Next == nil && carry > 0 {
			ptr.Next = &ListNode{Val: carry}
			break
		}
		ptr = ptr.Next
	}
	return ok(newhead)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func main() {
	// Example usage
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 8, Next: &ListNode{Val: 9}}}
	printList(head1)
	result1 := doubleIt(head1)
	printList(result1)

	head2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	printList(head2)
	result2 := doubleIt(head2)
	printList(result2)
}
