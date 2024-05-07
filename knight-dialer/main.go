package main

import "fmt"

var (
	dx  = []int{1, 1, -1, -1, 2, 2, -2, -2}
	dy  = []int{2, -2, 2, -2, 1, -1, 1, -1}
	MOD = 1_000_000_007
)

func ok(x, y int) bool {
	if (x == 3 && (y == 0 || y == 2)) || x < 0 || x > 3 || y < 0 || y > 2 {
		return false
	}
	return true
}

func f(x, y, n int, dp [][][]int) int {
	if !ok(x, y) {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if dp[x][y][n] != -1 {
		return dp[x][y][n]
	}

	cnt := 0
	for k := 0; k < 8; k++ {
		i, j := x+dx[k], y+dy[k]
		if ok(i, j) {
			cnt = (cnt%MOD + f(i, j, n-1, dp)%MOD) % MOD
		}
	}

	dp[x][y][n] = cnt % MOD
	return dp[x][y][n]
}

func knightDialer(n int) int {
	cnt := 0
	dp := make([][][]int, 4)
	for i := range dp {
		dp[i] = make([][]int, 3)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			cnt = (cnt%MOD + f(i, j, n, dp)%MOD) % MOD
		}
	}
	return cnt % MOD
}

func main() {
	fmt.Println(knightDialer(1))
	fmt.Println(knightDialer(2))
	//fmt.Println(knightDialer())
}
