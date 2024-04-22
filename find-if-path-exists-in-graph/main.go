package main

func dfs(curr int, g [][]int, vis *[]bool) {
	for _, ele := range g[curr] {
		if !(*vis)[ele] {
			(*vis)[ele] = true
			dfs(ele, g, vis)
		}
	}
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	g := make([][]int, n)
	vis := make([]bool, n)
	for I := range edges {
		u, v := edges[I][0], edges[I][1]
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	vis[source] = true
	dfs(source, g, &vis)
	return vis[destination] == true
}

func main() {

}
