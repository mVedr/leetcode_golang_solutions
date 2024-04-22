package main

import "fmt"

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

type Cell struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	q := []Cell{}
	tm := 0
	N, M := len(grid), len(grid[0])

	for I := range grid {
		for J := range grid[I] {
			if grid[I][J] == 2 {
				q = append(q, Cell{x: I, y: J})
			}
		}
	}
	for len(q) > 0 {
		n := len(q)
		tm++
		for I := 0; I < n; I++ {
			xx, yy := q[0].x, q[0].y
			q = q[1:]
			for K := range 4 {
				cx, cy := xx+dx[K], yy+dy[K]
				if cx >= 0 && cx < N && cy >= 0 && cy < M && grid[cx][cy] == 1 {
					grid[cx][cy] = 2
					q = append(q, Cell{x: cx, y: cy})
				}
			}
		}
	}
	for I := range N {
		for J := range M {
			if grid[I][J] == 1 {
				return -1
			}
		}
	}
	if tm == 0 {
		return 0
	}
	return tm - 1
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))
}
