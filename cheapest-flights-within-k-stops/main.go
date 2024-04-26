package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	k, cost, node int
}

type Heap []Node

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Less(i int, j int) bool {
	return h[i].k < h[j].k

}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h Heap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, kk int) int {
	h := &Heap{}
	heap.Init(h)
	g := make([][][2]int, n)
	for _, f := range flights {
		u, v, w := f[0], f[1], f[2]
		g[u] = append(g[u], [2]int{v, w})
	}
	cc := make([]int, n)
	for I := range cc {
		cc[I] = math.MaxInt32
	}
	cc[src] = 0
	heap.Push(h, Node{k: 0, cost: 0, node: src})
	for h.Len() > 0 {
		tp := heap.Pop(h).(Node)
		if tp.k <= kk {
			for _, neigh := range g[tp.node] {
				if cc[neigh[0]] > tp.cost+neigh[1] {
					cc[neigh[0]] = tp.cost + neigh[1]
					heap.Push(h, Node{node: neigh[0], cost: cc[neigh[0]], k: tp.k + 1})
				}
			}
		}
	}
	if cc[dst] == math.MaxInt32 {
		return -1
	}
	return cc[dst]
}

func main() {
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1))
	fmt.Println(findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0))
}
