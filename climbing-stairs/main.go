package main

import "fmt"

func solve(n int, dp []int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if dp[n] != -1 {
		return dp[n]
	}
	ans := 0
	if n >= 1 {
		ans += solve(n-1, dp)
	}
	if n >= 2 {
		ans += solve(n-2, dp)
	}
	dp[n] = ans
	return ans
}

func main() {
	dp := make([]int, 51)
	for i := range dp {
		dp[i] = -1
	}
	solve(45, dp)
	fmt.Println(dp[2])
	fmt.Println(dp[3])
}
