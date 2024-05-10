package main

import (
	"container/heap"
)

type Node struct {
	val  float64
	a, b int
}

type Heap []Node

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	h := &Heap{}
	heap.Init(h)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			node := Node{val: float64(arr[i]) / float64(arr[j]), a: arr[i], b: arr[j]}
			heap.Push(h, node)
		}
	}

	for i := 1; i < k; i++ {
		heap.Pop(h)
	}

	result := heap.Pop(h).(Node)
	return []int{result.a, result.b}
}

func main() {

}
