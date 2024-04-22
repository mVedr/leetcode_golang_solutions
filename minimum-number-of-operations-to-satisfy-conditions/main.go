package main

import (
	"fmt"
	"math"
)

func calc(num, col int, grid [][]int) int {
	ans := 0
	for I := range len(grid) {
		if grid[I][col] != num {
			ans++
		}
	}
	return ans
}

func solve(i, prev, n int, grid [][]int, dp *[][]int) int {
	if i >= n {
		return 0
	}
	if (*dp)[i][prev+1] != -1 {
		return (*dp)[i][prev+1]
	}
	ans := math.MaxInt32
	for I := range 10 {
		if I != prev {
			ans = min(ans, calc(I, i, grid)+solve(i+1, I, n, grid, dp))
		}
	}
	(*dp)[i][prev+1] = ans
	return ans
}

func Init2d[T any](n, m int, val T) [][]T {
	arr := make([][]T, n)
	for i := range n {
		arr[i] = make([]T, m)
	}
	for i := range n {
		for j := range m {
			arr[i][j] = val
		}
	}
	return arr
}

func minimumOperations(grid [][]int) int {
	dp := Init2d(1005, 15, -1)
	return solve(0, -1, len(grid[0]), grid, &dp)
}

func main() {
	fmt.Println(minimumOperations([][]int{{1, 0, 2}, {1, 0, 2}}))
	fmt.Println(minimumOperations([][]int{{1, 1, 1}, {0, 0, 0}}))
	fmt.Println(minimumOperations([][]int{{1}, {2}, {3}}))
}
