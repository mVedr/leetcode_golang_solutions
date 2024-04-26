package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	node int
	dist int
}

type Heap []Node

func (h *Heap) Len() int {
	return len(*h)
}

func (h Heap) Less(i int, j int) bool {
	return h[i].dist < h[j].dist
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

func minimumTime(n int, edges [][]int, disappear []int) []int {
	g := make([][]Node, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], Node{node: e[1], dist: e[2]})
		g[e[1]] = append(g[e[1]], Node{node: e[0], dist: e[2]})
	}
	h := &Heap{}
	heap.Init(h)
	dis := make([]int, n)
	for I := range dis {
		dis[I] = math.MaxInt32
	}
	dis[0] = 0
	heap.Push(h, Node{node: 0, dist: 0})
	for h.Len() > 0 {
		tp := heap.Pop(h).(Node)
		for _, neigh := range g[tp.node] {
			if dis[neigh.node] > dis[tp.node]+neigh.dist && disappear[neigh.node] > dis[tp.node]+neigh.dist {
				dis[neigh.node] = dis[tp.node] + neigh.dist
				heap.Push(h, Node{node: neigh.node, dist: dis[neigh.node]})
			}
		}
	}
	ans := []int{}
	for I, val := range disappear {
		if dis[I] > val {
			ans = append(ans, -1)
		} else {
			ans = append(ans, dis[I])
		}
	}
	return ans
}

func main() {
	fmt.Println(minimumTime(3, [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}, []int{0, -1, 4}))
	fmt.Println(minimumTime(3, [][]int{{0, 1, 2}, {1, 2, 1}, {0, 2, 4}}, []int{1, 3, 5}))
}
