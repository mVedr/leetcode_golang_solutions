package main

import (
	"container/heap"
	"math"
	"sort"
)

type Node struct {
	quality, wage int
}

type Heap []int

func (pq Heap) Len() int {
	return len(pq)
}
func (pq Heap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq Heap) Less(i, j int) bool {
	return pq[i] > pq[j]
}
func (pq *Heap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}
func (pq *Heap) Pop() interface{} {
	x := (*pq)[pq.Len()-1]
	*pq = (*pq)[0 : pq.Len()-1]
	return x
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	w := make([]Node, len(quality))
	for i := range w {
		w[i] = Node{quality[i], wage[i]}
	}

	sort.Slice(w, func(i, j int) bool {
		return w[i].wage*w[j].quality < w[j].wage*w[i].quality
	})

	pq, qSum, minCost := make(Heap, 0), 0, math.MaxFloat64
	for _, w := range w {
		heap.Push(&pq, w.quality)
		qSum += w.quality
		if pq.Len() > k {
			qSum -= heap.Pop(&pq).(int)
		}
		cost := float64(w.wage*qSum) / float64(w.quality)
		if pq.Len() == k && cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func main() {

}
