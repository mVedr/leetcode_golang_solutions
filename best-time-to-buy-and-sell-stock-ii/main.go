package main

import "fmt"

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

func solve(i int, b int, n int, prices []int, dp *[][]int) int {
	if i >= n {
		return 0
	}
	if (*dp)[i][b] != -1 {
		return (*dp)[i][b]
	}
	ans := 0
	if b == 0 {
		ans = max(ans, -prices[i]+solve(i+1, b+1, n, prices, dp))
		ans = max(ans, solve(i+1, b, n, prices, dp))
	} else {
		ans = max(ans, prices[i]+solve(i+1, b-1, n, prices, dp))
		ans = max(ans, solve(i+1, b, n, prices, dp))
	}
	(*dp)[i][b] = ans
	return ans
}

func maxProfit(prices []int) int {
	dp := Init2d(40002, 3, -1)
	return solve(0, 0, len(prices), prices, &dp)
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
