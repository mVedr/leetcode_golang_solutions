package main

import (
	"fmt"
	"math"
)

func solve(cx, py int, grid [][]int, dp *map[[2]int]int) int {
	if cx >= len(grid) {
		return 0
	}
	if val, ok := (*dp)[[2]int{cx, py}]; ok {
		return val
	}
	ans := math.MaxInt32
	for I := range len(grid[0]) {
		if I != py {
			ans = min(ans, grid[cx][I]+solve(cx+1, I, grid, dp))
		}
	}
	(*dp)[[2]int{cx, py}] = ans
	return ans
}

func minFallingPathSum(grid [][]int) int {
	dp := make(map[[2]int]int)
	return solve(0, -1, grid, &dp)
}

func main() {
	fmt.Println(minFallingPathSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(minFallingPathSum([][]int{{7}}))
}
