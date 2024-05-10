package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	ca, cb := 2, 0
	pa, pb := list1, list1
	for ca <= a {
		pa = pa.Next
		ca++
	}
	for cb <= b {
		pb = pb.Next
		cb++
	}
	//fmt.Println(pa.Val, " ", pb.Val)
	pa.Next = list2
	last2 := list2
	for last2.Next != nil {
		last2 = last2.Next
	}
	last2.Next = pb
	return list1
}

func main() {

}
