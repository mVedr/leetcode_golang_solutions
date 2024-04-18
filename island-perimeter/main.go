package main

import "fmt"

var (
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}
)

func calculate(cx, cy int, g [][]int) int {
	n, m := len(g), len(g[0])
	ans := 0
	for I := range 4 {
		nx, ny := cx+dx[I], cy+dy[I]
		if nx >= 0 && nx < n && ny >= 0 && ny < m {
			if g[nx][ny] == 0 {
				ans++
			}
		} else {
			ans++
		}
	}
	return ans
}

func solve(g [][]int) int {
	ans := 0

	for I := range g {
		for J := range g[0] {
			if g[I][J] == 1 {
				ans += calculate(I, J, g)
			}

		}
	}
	return ans
}

func main() {
	fmt.Println(solve([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	fmt.Println(solve([][]int{{1, 0}}))
}
