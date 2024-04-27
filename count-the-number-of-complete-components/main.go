package main

func dfs(curr int, nc, ec *int, g [][]int, vis *[]bool) {
	(*vis)[curr] = true
	(*nc)++
	(*ec) += len(g[curr])
	for _, neigh := range g[curr] {
		if !(*vis)[neigh] {
			dfs(neigh, nc, ec, g, vis)
		}
	}
}

func countCompleteComponents(n int, edges [][]int) int {
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	vis := make([]bool, n)
	ans := 0
	for I := range n {
		nc, ec := 0, 0
		if !vis[I] {
			ec := 0
			nc := 0
			dfs(I, &nc, &ec, g, &vis)
		}
		if nc*(nc-1) == ec {
			ans++
		}
	}
	return ans
}

func main() {

}
