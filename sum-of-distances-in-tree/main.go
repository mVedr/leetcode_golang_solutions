package main

import "fmt"

func dfs(node, pr int, par, num *[]int, vis *[]bool, g [][]int) int {
	(*par)[node] = pr
	ans := 0
	for _, neigh := range g[node] {
		if !(*vis)[neigh] {
			(*vis)[neigh] = true
			ans += (1 + dfs(neigh, node, par, num, vis, g))
			ans += (*num)[neigh]
		}
	}
	(*num)[node] = ans
	return 0
}

func bfs(g [][]int) int {
	vis := make([]bool, len(g))
	ans := 0
	q := []int{}
	q = append(q, 0)
	vis[0] = true
	lvl := 0
	for len(q) > 0 {
		n := len(q)
		lvl++
		for range n {
			tp := q[0]
			q = q[1:]
			for _, neigh := range g[tp] {
				if !vis[neigh] {
					ans += lvl
					vis[neigh] = true
					q = append(q, neigh)
				}
			}
		}
	}
	return ans
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	par, num := make([]int, n), make([]int, n)
	vis := make([]bool, n)
	//dist := make([]int,n)
	for I := range par {
		num[I] = 0
		par[I] = -1
	}
	vis[0] = true
	dfs(0, -1, &par, &num, &vis, g)

	ans1 := bfs(g)
	ans := make([]int, n)

	Q := []int{}
	V := make([]bool, n)
	V[0] = true
	ans[0] = ans1
	Q = append(Q, 0)
	for len(Q) > 0 {
		N := len(Q)
		for range N {
			tp := Q[0]
			Q = Q[1:]
			if tp > 0 {
				ans[tp] = ans[par[tp]] + n - 2 - 2*num[tp]
			}
			for _, neigh := range g[tp] {
				if !V[neigh] {
					V[neigh] = true
					Q = append(Q, neigh)
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(sumOfDistancesInTree(6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}))
	fmt.Println(sumOfDistancesInTree(1, [][]int{}))
	fmt.Println(sumOfDistancesInTree(2, [][]int{{1, 0}}))
	fmt.Println(sumOfDistancesInTree(3, [][]int{{2, 1}, {0, 2}}))
}
