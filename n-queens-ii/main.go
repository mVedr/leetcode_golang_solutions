package main

import "bytes"

func ok(x, y, n int, vis [][]int) bool {
	if y == 0 {
		return true
	}

	for i := 0; i < y; i++ {
		if vis[x][i] == 1 {
			return false
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if vis[i][j] == 1 {
			return false
		}
	}

	for i, j := x+1, y-1; i < n && j >= 0; i, j = i+1, j-1 {
		if vis[i][j] == 1 {
			return false
		}
	}

	return true
}

func solve(y int, n int, vis [][]int, ans *int) {
	if y == n {
		temp := []string{}
		for _, row := range vis {
			buf := new(bytes.Buffer)
			for _, cell := range row {
				if cell == 1 {
					buf.WriteByte('Q')
				} else {
					buf.WriteByte('.')
				}
			}
			temp = append(temp, buf.String())
		}
		*ans = (*ans) + 1
		return
	}

	for x := 0; x < n; x++ {
		if ok(x, y, n, vis) {
			vis[x][y] = 1
			solve(y+1, n, vis, ans)
			vis[x][y] = 0
		}
	}
}

func Init2d[T any](n, m int, val T) [][]T {
	arr := make([][]T, n)
	for i := range arr {
		arr[i] = make([]T, m)
		for j := range arr[i] {
			arr[i][j] = val
		}
	}
	return arr
}

func totalNQueens(n int) int {
	ans := 0
	vis := Init2d(n, n, 0)
	solve(0, n, vis, &ans)
	return ans
}

func main() {

}
