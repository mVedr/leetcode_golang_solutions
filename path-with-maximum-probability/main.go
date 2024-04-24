package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	prob float64
	u    int
}

type NodeHeap []Node

func (n NodeHeap) Len() int {
	return len(n)
}

func (n NodeHeap) Less(i int, j int) bool {
	return n[i].prob > n[j].prob
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(Node))
}

func (h NodeHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {
	h := &NodeHeap{}
	heap.Init(h)
	dist := make([]float64, n)
	for i := range dist {
		dist[i] = -1.0
	}
	g := make([][]Node, n)

	for ind, E := range edges {
		g[E[0]] = append(g[E[0]], Node{u: E[1], prob: succProb[ind]})
		g[E[1]] = append(g[E[1]], Node{u: E[0], prob: succProb[ind]})
	}

	dist[start_node] = 1.0
	heap.Push(h, Node{u: start_node, prob: dist[start_node]})

	for h.Len() > 0 {
		tp := heap.Pop(h).(Node)
		for _, neigh := range g[tp.u] {
			if dist[neigh.u] < neigh.prob*dist[tp.u] {
				dist[neigh.u] = neigh.prob * dist[tp.u]
				heap.Push(h, Node{u: neigh.u, prob: dist[neigh.u]})
			}
		}
	}
	if dist[end_node] == -1 {
		return 0.0
	}
	return dist[end_node]
}

func main() {
	fmt.Println(maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.2}, 0, 2))
	fmt.Println(maxProbability(3, [][]int{{0, 1}, {1, 2}, {0, 2}}, []float64{0.5, 0.5, 0.3}, 0, 2))
	fmt.Println(maxProbability(3, [][]int{{0, 1}}, []float64{0.5}, 0, 2))
}
