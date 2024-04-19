package main

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

func ok(cx, cy, n, m int) bool {
	return cx >= 0 && cx < n && cy >= 0 && cy < m
}

func dfs(cx, cy, n, m int, grid [][]byte, vis *[][]int) {
	for i := range dx {
		xx, yy := cx+dx[i], cy+dy[i]
		if ok(xx, yy, n, m) && (*vis)[xx][yy] == 0 && grid[xx][yy] == '1' {
			(*vis)[xx][yy] = 1
			dfs(xx, yy, n, m, grid, vis)
		}
	}
}

func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	vis := make([][]int, n)
	for i := range vis {
		vis[i] = make([]int, m)
	}
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if vis[i][j] == 0 && grid[i][j] == '1' {
				ans++
				vis[i][j] = 1
				dfs(i, j, n, m, grid, &vis)
			}
		}
	}
	return ans
}

func main() {

}
