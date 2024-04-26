package main

import "fmt"

func solve(curr int, color *[]int, g [][]int) bool {
	ans := true
	for _, neigh := range g[curr] {
		if (*color)[neigh] == -1 {
			(*color)[neigh] = 1 - (*color)[curr]
			ans = ans && solve(neigh, color, g)
		} else if (*color)[neigh] == (*color)[curr] {
			return false
		}
	}
	return ans
}

func isBipartite(g [][]int) bool {
	color := make([]int, len(g))
	for I := range color {
		color[I] = -1
	}
	ans := true
	for I := range g {
		if color[I] == -1 {
			color[I] = 0
			ans = ans && solve(I, &color, g)
		}
	}
	return ans
}

func main() {
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}))
}
