package main

import "fmt"

const MOD = 1000000007

func dfs(prev, remz, remo, limit int, dp [][][]int) int {
	if remz == 0 && remo == 0 {
		return 1
	} else if remz < 0 || remo < 0 {
		return 0
	}
	if dp[prev][remz][remo] != -1 {
		return dp[prev][remz][remo]
	}
	ans := 0
	if prev == 1 {
		for i := 1; i <= limit; i++ {
			ans += dfs(0, remz-i, remo, limit, dp)
			ans %= MOD
		}
	} else {
		for i := 1; i <= limit; i++ {
			ans += dfs(1, remz, remo-i, limit, dp)
			ans %= MOD
		}
	}
	dp[prev][remz][remo] = ans
	return ans
}

func numberOfStableArrays(remz, remo, limit int) int {
	dp := make([][][]int, 2)
	for i := range dp {
		dp[i] = make([][]int, remz+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, remo+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	res := 0
	for i := 1; i <= limit; i++ {
		res += dfs(0, remz-i, remo, limit, dp)
		res %= MOD
	}
	for i := 1; i <= limit; i++ {
		res += dfs(1, remz, remo-i, limit, dp)
		res %= MOD
	}
	return res
}

func main() {
	fmt.Println(numberOfStableArrays(1, 1, 2))
	fmt.Println(numberOfStableArrays(1, 2, 1))
	fmt.Println(numberOfStableArrays(3, 3, 2))
}
