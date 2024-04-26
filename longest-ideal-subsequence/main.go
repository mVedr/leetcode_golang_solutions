package main

import "fmt"

func ab(a, b int) int {
	if a-b >= 0 {
		return a - b
	}
	return b - a
}

func solve(i, prev, k int, s string, dp *[][]int) int {
	if i >= len(s) {
		return 0
	}
	if (*dp)[i][prev+1] != -1 {
		return (*dp)[i][prev+1]
	}
	ans := 0
	if prev == -1 {
		ans = max(ans, 1+solve(i+1, int(s[i]-'a'), k, s, dp))
	} else {
		diff := ab(int(s[i]-'a'), prev)
		if diff <= k {
			ans = max(ans, 1+solve(i+1, int(s[i]-'a'), k, s, dp))
		}
	}

	ans = max(ans, 0+solve(i+1, prev, k, s, dp))
	(*dp)[i][prev+1] = ans
	return ans
}

func Init2d[T any](n, m int, val T) [][]T {
	arr := make([][]T, n)
	for i := range n {
		arr[i] = make([]T, m)
	}
	for i := range n {
		for j := range m {
			arr[i][j] = val
		}
	}
	return arr
}

func longestIdealString(s string, k int) int {
	dp := Init2d(len(s)+1, 30, -1)

	return solve(0, -1, k, s, &dp)
}

func main() {
	fmt.Println(longestIdealString("acfgbd", 2))
	fmt.Println(longestIdealString("abcd", 3))
}
