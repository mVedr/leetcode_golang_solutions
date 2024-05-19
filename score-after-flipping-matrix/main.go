package main

func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := (1 << (n - 1)) * m
	for J := 1; J < n; J++ {
		cur := 0
		for I := 0; I < m; I++ {
			if grid[I][J] == grid[I][0] {
				cur++
			}
		}
		res += max(cur, m-cur) * (1 << (n - J - 1))
	}
	return res
}

func main() {

}
