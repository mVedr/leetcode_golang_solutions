package main

import "fmt"

func ok(x, y, n, m int) bool {
	return x >= 0 && x < n && y >= 0 && y < m
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

func numberOfRightTriangles(grid [][]int) int64 {
	ans := int64(0)
	numsR, numsC := make([]int, len(grid)), make([]int, len(grid[0]))
	for I := range grid {
		for J := range grid[0] {
			numsR[I] += grid[I][J]
		}
	}
	for J := range grid[0] {
		for I := range grid {
			numsC[J] += grid[I][J]
		}
	}
	for I := range grid {
		for J := range grid[0] {
			if grid[I][J] == 1 {
				uu := int64((numsR[I] - 1) * (numsC[J] - 1))
				if uu > 0 {
					ans += uu
				}

			}
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfRightTriangles([][]int{{0, 1, 0}, {0, 1, 1}, {0, 1, 0}}))
	fmt.Println(numberOfRightTriangles([][]int{{1, 0, 0, 0}, {0, 1, 0, 1}, {1, 0, 0, 0}}))
	fmt.Println(numberOfRightTriangles([][]int{{1, 0, 1}, {1, 0, 0}, {1, 0, 0}}))
}
