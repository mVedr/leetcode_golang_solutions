package main

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

func dfs(x, y int, grid *[][]int) {
	if x < 0 || x >= len(*grid) || y < 0 || y >= len((*grid)[0]) || (*grid)[x][y] == 0 {
		return
	}
	(*grid)[x][y] = 0
	for I := range 4 {
		dfs(x+dx[I], y+dy[I], grid)
	}
}

func numEnclaves(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	for I := range m {
		if grid[0][I] == 1 {
			dfs(0, I, &grid)
		}
		if grid[n-1][I] == 1 {
			dfs(n-1, I, &grid)
		}
	}
	for I := range n {
		if grid[I][0] == 1 {
			dfs(I, 0, &grid)
		}
		if grid[I][m-1] == 1 {
			dfs(I, m-1, &grid)
		}
	}
	ans := 0
	for I := range n {
		for J := range m {
			if grid[I][J] == 1 {
				ans++
			}
		}
	}
	return ans
}

func main() {

}
