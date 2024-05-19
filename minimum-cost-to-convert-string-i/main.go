package main

import (
	"container/heap"
)

var (
	INF = 1_000_000_000
)

type Node struct {
	u    byte
	dist int
}

type Heap []Node

// Len implements heap.Interface.
func (h *Heap) Len() int {
	return len(*h)
}

// Less implements heap.Interface.
func (h Heap) Less(i int, j int) bool {
	return h[i].dist < h[j].dist
}

// Pop implements heap.Interface.
func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Push implements heap.Interface.
func (h *Heap) Push(x any) {
	*h = append(*h, x.(Node))
}

// Swap implements heap.Interface.
func (h Heap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func sssp(src byte, g map[int][]Node) []int {
	dd := make([]int, 26)
	for I := range dd {
		dd[I] = INF
	}
	h := &Heap{}
	heap.Init(h)
	dd[src-'a'] = 0
	h.Push(Node{u: src, dist: 0})
	for h.Len() > 0 {
		tp := h.Pop().(Node)
		tpd := tp.dist
		no := int(tp.u - 'a')
		for _, neigh := range g[no] {
			if dd[neigh.u-'a'] > tpd+neigh.dist {
				dd[neigh.u-'a'] = tpd + neigh.dist
				h.Push(Node{u: neigh.u, dist: tpd + neigh.dist})
			}
		}
	}
	return dd
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {

	g := map[int][]Node{}
	for I := range original {
		g[int(original[I]-'a')] = append(g[int(original[I]-'a')], Node{u: changed[I], dist: cost[I]})
	}
	ans := int64(0)
	ss := [26][26]int{}
	for I := range 26 {
		ss[I] = [26]int(sssp(byte('a'+I), g))
	}
	for I := range source {
		u, v := int(source[I]-'a'), int(target[I]-'a')
		if ss[u][v] == INF {
			return -1
		}
		ans += int64(ss[u][v])
	}
	return ans
}

func main() {

}
