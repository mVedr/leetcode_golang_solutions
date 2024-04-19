package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Comp struct {
	Node *ListNode
}

type NodeHeap []Comp

func (n *NodeHeap) Len() int {
	return len(*n)
}

func (n *NodeHeap) Less(i int, j int) bool {
	return (*n)[i].Node.Val < (*n)[j].Node.Val
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (n *NodeHeap) Push(x any) {
	*n = append(*n, x.(Comp))
}

func (n *NodeHeap) Swap(i int, j int) {
	(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &NodeHeap{}
	heap.Init(h)
	dum := &ListNode{Val: -1}
	ptr := dum
	for _, node := range lists {
		if node != nil {
			heap.Push(h, &Comp{Node: node})
		}
	}

	for h.Len() > 0 {
		cmp := (*h)[0]
		heap.Pop(h)
		if cmp.Node.Next != nil {
			heap.Push(h, &Comp{Node: cmp.Node.Next})
		}
		ptr.Next = cmp.Node
		ptr = ptr.Next
	}

	return dum.Next
}

func main() {

}
