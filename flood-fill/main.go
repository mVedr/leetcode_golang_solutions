package main

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

type Cell struct {
	x int
	y int
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

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	N, M := len(image), len(image[0])
	q := []Cell{{x: sr, y: sc}}
	st := image[sr][sc]
	image[sr][sc] = color

	for len(q) > 0 {
		n := len(q)
		for _ = range n {
			cx, cy := q[0].x, q[0].y
			q = q[1:]
			for I := range 4 {
				xx, yy := cx+dx[I], cy+dy[I]
				if xx >= 0 && xx < N && yy >= 0 && yy < M && image[xx][yy] == st {
					image[xx][yy] = color
					q = append(q, Cell{x: xx, y: yy})
				}
			}
		}
	}
	return image
}

func main() {

}
