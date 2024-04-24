package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {
	g := make([][]int, n)
	for i := range n {
		g[i] = []int{}
	}
	indeg := make([]int, n)
	for _, E := range edges {
		g[E[0]] = append(g[E[0]], E[1])
		g[E[1]] = append(g[E[1]], E[0])
		indeg[E[0]]++
		indeg[E[1]]++
	}
	q := []int{}
	for i := range indeg {
		if indeg[i] == 1 {
			indeg[i]--
			q = append(q, i)
		}
	}
	ans := []int{}
	for len(q) > 0 {
		N := len(q)
		ans = []int{}
		for _ = range N {
			tp := q[0]
			q = q[1:]
			ans = append(ans, tp)
			for _, neigh := range g[tp] {
				indeg[neigh]--
				if indeg[neigh] == 1 {
					q = append(q, neigh)
				}

			}
		}
	}
	if n == 1 {
		return []int{}
	}
	return ans
}

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}}))
	fmt.Println(findMinHeightTrees(3, [][]int{{0, 1}, {0, 2}}))
}
