package main

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
	A  = 0
)

func solve(i, j int, ans int, vis [][]bool, grid [][]int) {
	for I := range 4 {
		cx, cy := i+dx[I], j+dy[I]
		if cx >= 0 && cx < len(vis) && cy >= 0 && cy < len(vis[0]) {
			if !vis[cx][cy] {
				vis[cx][cy] = true
				A = max(A, ans+grid[cx][cy])
				solve(cx, cy, ans+grid[cx][cy], vis, grid)
				vis[cx][cy] = false
			}

		}
	}
}

func getMaximumGold(grid [][]int) int {
	vis := make([][]bool, len(grid))
	for I := range vis {
		vis[I] = make([]bool, len(grid[0]))
	}
	for I := range grid {
		for J := range grid[0] {
			if grid[I][J] != 0 {
				vis[I][J] = true
				solve(I, J, grid[I][J], vis, grid)
				vis[I][J] = false
			}
		}
	}
	return A
}

func main() {

}
