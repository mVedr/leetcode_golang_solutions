package main

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

func reverseStr(str string) string {
	strSlice := []rune(str)
	l := len(strSlice)

	for i := 0; i < (l / 2); i++ {
		strSlice[i], strSlice[l-i-1] = strSlice[l-i-1], strSlice[i]
	}
	return string(strSlice)
}

func longestPalindromeSubseq(s string) int {
	return longestCommonSubsequence(s, reverseStr(s))
}

func minInsertions(s string) int {
	return len(s) - longestPalindromeSubseq(s)
}

func main() {

}
