package main

import "fmt"

func solve(i, j, m, n int, dp *[][]int) int {
	if i == m-1 && j == n-1 {
		return 1
	}
	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}
	ans := 0
	if i+1 < m {
		ans += solve(i+1, j, m, n, dp)
	}
	if j+1 < n {
		ans += solve(i, j+1, m, n, dp)
	}
	(*dp)[i][j] = ans
	return ans
}

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return solve(0, 0, m, n, &dp)
}

func main() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(3, 2))
}
