package main

import "fmt"

func findChampion(n int, edges [][]int) int {
	if n == 1 {
		return 0
	}
	g := make([][]int, n)
	indeg, vis := make([]int, n), make([]int, n)
	for _, e := range edges {
		g[e[1]] = append(g[e[1]], e[0])
		indeg[e[0]]++
		vis[e[0]] = 1
		vis[e[1]] = 1
	}
	ans := []int{}
	q := []int{}
	for i := range n {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		n := len(q)

		for range n {
			node := q[0]
			if len(g[node]) == 0 {
				ans = append(ans, node)
			}

			q = q[1:]
			for _, neigh := range g[node] {
				if indeg[neigh] == 1 {
					q = append(q, neigh)
				}
				indeg[neigh]--
			}
		}
	}
	if len(ans) != 1 {
		return -1
	}
	for _, v := range vis {
		if v == 0 {
			return -1
		}
	}
	return ans[0]
}

func main() {
	fmt.Println(findChampion(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findChampion(4, [][]int{{0, 2}, {1, 3}, {1, 2}}))
	fmt.Println(findChampion(3, [][]int{{0, 1}}))
}
