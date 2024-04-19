package main

import "fmt"

func solve(i, j int, s, t string, dp *[][]int) int {
	if i >= len(s) || j >= len(t) {
		return 0
	}
	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}
	ans := 0
	if s[i] == t[j] {
		ans = max(ans, 1+solve(i+1, j+1, s, t, dp))
	} else {
		ans = max(ans, solve(i, j+1, s, t, dp))
		ans = max(ans, solve(i+1, j, s, t, dp))
	}
	(*dp)[i][j] = ans
	return ans
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := range dp {
		dp[i] = make([]int, len(text2))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return solve(0, 0, text1, text2, &dp)
}

func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))
}
