package main

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		res[i] = make([]int, n-2)
	}
	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			res[i-1][j-1] = max(
				grid[i-1][j-1],
				grid[i-1][j],
				grid[i-1][j+1],
				grid[i][j-1],
				grid[i][j],
				grid[i][j+1],
				grid[i+1][j-1],
				grid[i+1][j],
				grid[i+1][j+1],
			)
		}
	}
	return res
}

func main() {

}
