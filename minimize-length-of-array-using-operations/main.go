package main

import (
	"fmt"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumArrayLength(nums []int) int {
	mn := math.MaxInt
	for _, e := range nums {
		mn = min(e, mn)
	}
	cc := 0
	for _, e := range nums {
		if mn == e {
			cc++
		}
		if e%mn != 0 {
			return 1
		}
	}
	return (cc + 1) / 2
}
func main() {
	fmt.Println(minimumArrayLength([]int{1, 4, 3, 1}))
	fmt.Println(minimumArrayLength([]int{5, 5, 5, 10, 5}))
	fmt.Println(minimumArrayLength([]int{2, 3, 4}))
	fmt.Println(minimumArrayLength([]int{4, 4, 4}))
}
