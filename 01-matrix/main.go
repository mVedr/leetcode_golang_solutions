package main

import "fmt"

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

type Cell struct {
	x, y, dist int
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

func updateMatrix(mat [][]int) [][]int {
	q := []Cell{}
	dist := Init2d(len(mat), len(mat[0]), -1)
	N, M := len(mat), len(mat[0])
	for i := range len(mat) {
		for j := range len(mat[0]) {
			if mat[i][j] == 0 {
				q = append(q, Cell{x: i, y: j, dist: 0})
				dist[i][j] = 0
			}
		}
	}

	for len(q) > 0 {
		n := len(q)
		for _ = range n {
			tp := q[0]
			q = q[1:]
			cx, cy, d := tp.x, tp.y, tp.dist
			for I := range 4 {
				xx, yy := cx+dx[I], cy+dy[I]
				if xx >= 0 && xx < N && yy >= 0 && yy < M && dist[xx][yy] == -1 {
					dist[xx][yy] = d + 1
					q = append(q, Cell{x: xx, y: yy, dist: d + 1})
				}
			}

		}
	}

	return dist
}

func main() {
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}))
}
