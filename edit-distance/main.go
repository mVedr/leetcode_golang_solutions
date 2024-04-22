package main

import (
	"fmt"
	"math"
)

func solve(i, j int, s, t string, dp *[][]int) int {
	if i < 0 && j < 0 {
		return 0
	}
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}
	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}
	ans := math.MaxInt32
	if s[i] == t[j] {
		ans = min(ans, solve(i-1, j-1, s, t, dp))
	} else {
		ans = min(ans, 1+solve(i, j-1, s, t, dp))
		ans = min(ans, 1+solve(i-1, j, s, t, dp))
		ans = min(ans, 1+solve(i-1, j-1, s, t, dp))
	}
	(*dp)[i][j] = ans
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

func minDistance(s string, t string) int {
	dp := Init2d(505, 505, -1)
	return solve(len(s)-1, len(t)-1, s, t, &dp)
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
}
