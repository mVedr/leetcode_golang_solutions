package main

import "fmt"

func solve(b [][]byte) {
	if len(b) < 2 || len(b[0]) < 2 {
		return
	}

	for i := 0; i < len(b); i++ {
		dfs(b, i, 0)
		dfs(b, i, len(b[i])-1)
	}
	for j := 0; j < len(b[0]); j++ {
		dfs(b, 0, j)
		dfs(b, len(b)-1, j)
	}
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == 'Z' {
				b[i][j] = 'O'
			} else {
				b[i][j] = 'X'
			}
		}
	}
	fmt.Println(b)
}

func dfs(b [][]byte, i, j int) {
	if i < 0 || i >= len(b) || j < 0 || j >= len(b[i]) ||
		b[i][j] != 'O' {
		return
	}
	b[i][j] = 'Z'

	if i-1 >= 0 {
		dfs(b, i-1, j)
	}

	if j+1 < len(b[i]) {
		dfs(b, i, j+1)
	}

	if i+1 < len(b) {
		dfs(b, i+1, j)
	}

	if j-1 >= 0 {
		dfs(b, i, j-1)
	}
}

func main() {
	b1 := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	b2 := [][]byte{{'X'}}
	solve(b1)
	solve(b2)
}
