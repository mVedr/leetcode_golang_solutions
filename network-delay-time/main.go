package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	U    int
	Dist int
}

type NodeHeap []Node

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i int, j int) bool {
	return h[i].Dist < h[j].Dist
}

func (h NodeHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(a any) {
	*h = append(*h, a.(Node))
}

func (h *NodeHeap) Pop() any {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}

func networkDelayTime(times [][]int, n int, k int) int {
	dd := make([]int, n+1)
	adj := make(map[int][]Node)

	for i := 0; i < len(times); i++ {
		u1 := times[i][0]
		u2 := times[i][1]
		dist := times[i][2]

		adj[u1] = append(adj[u1], Node{U: u2, Dist: dist})
	}

	for i := 1; i <= n; i++ {
		dd[i] = math.MaxInt64
	}

	h := &NodeHeap{}
	heap.Push(h, Node{U: k, Dist: 0})
	dd[k] = 0

	for h.Len() > 0 {
		node := heap.Pop(h).(Node)
		v := node.U
		dist := node.Dist

		if dd[v] < dist {
			continue
		}

		for i := 0; i < len(adj[v]); i++ {
			nextNode := adj[v][i]
			neigh := nextNode.U
			nextTime := nextNode.Dist

			if dist+nextTime < dd[neigh] {
				dd[neigh] = dist + nextTime
				heap.Push(h, Node{U: neigh, Dist: dd[neigh]})
			}
		}
	}

	mx := 0
	for i := 0; i < len(dd); i++ {
		mx = max(mx, dd[i])
	}

	if mx == math.MaxInt64 {
		return -1
	}

	return mx
}

func main() {
	fmt.Println(networkDelayTime([][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2))
	fmt.Println(networkDelayTime([][]int{{1, 2, 1}}, 2, 1))
	fmt.Println(networkDelayTime([][]int{{1, 2, -1}}, 2, 1))
}
