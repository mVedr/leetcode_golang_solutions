package main

import "fmt"

func dfs(i int, g [][]int, vis []bool) {
	for j, ok := range g[i] {
		if ok == 1 && !vis[j] {
			vis[j] = true
			dfs(j, g, vis)
		}
	}
}

func solve(g [][]int) int {
	n := len(g)
	vis := make([]bool, n)
	ans := 0
	for i := range n {
		if !vis[i] {
			vis[i] = true
			ans++
			dfs(i, g, vis)
		}
	}
	return ans
}

func main() {
	g1 := [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}
	g2 := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

	fmt.Println(solve(g1))
	fmt.Println(solve(g2))
}
