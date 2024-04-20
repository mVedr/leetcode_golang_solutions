package main

import "fmt"

func findBottom(cx, cy int, land [][]int) (int, int) {
	n, m := len(land), len(land[0])
	for true {
		if cx+1 < n && cy < m && land[cx+1][cy] == 1 {
			cx++
		} else if cx < n && cy-1 >= 0 && land[cx][cy-1] == 1 {
			cy--
		} else {
			return cx, cy
		}
	}
	return cx, cy
}

func findUp(cx, cy int, land [][]int) (int, int) {
	n, m := len(land), len(land[0])
	for true {
		if cy+1 < m && cx < n && land[cx][cy+1] == 1 {
			cy++
		} else if cy < m && cx-1 >= 0 && land[cx-1][cy] == 1 {
			cx--
		} else {
			return cx, cy
		}
	}
	return cx, cy
}

func findFarmland(land [][]int) [][]int {
	vis := make([][]bool, len(land))
	for i := range vis {
		vis[i] = make([]bool, len(land[0]))
	}
	ans := [][]int{}
	for i := range land {
		for j := range land[i] {
			if land[i][j] == 1 {
				bx, by := findBottom(i, j, land)
				if !vis[bx][by] {
					vis[bx][by] = true
					tx, ty := findUp(bx, by, land)
					for I := tx; I <= bx; I++ {
						for J := by; J <= by; J++ {
							vis[I][J] = true
						}
					}
					ans = append(ans, []int{tx, by, bx, ty})
				}
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findFarmland([][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}))
	fmt.Println(findFarmland([][]int{{1, 1}, {1, 1}}))
	fmt.Println(findFarmland([][]int{{0}}))
}
