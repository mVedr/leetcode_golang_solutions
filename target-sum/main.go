package main

import "fmt"

func solve(i, sum, n, t int, nums []int, dp *[][]int) int {
	if i == n {
		if sum == t {
			return 1
		}
		return 0
	}
	if (*dp)[sum+20000][i] != -1 {
		return (*dp)[sum+20000][i]
	}
	ans := 0
	ans += solve(i+1, sum-nums[i], n, t, nums, dp)
	ans += solve(i+1, sum+nums[i], n, t, nums, dp)
	(*dp)[sum+20000][i] = ans
	return ans
}

func findTargetSumWays(nums []int, target int) int {
	dp := make([][]int, 40002)
	for i := range dp {
		dp[i] = make([]int, 22)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return solve(0, 0, len(nums), target, nums, &dp)
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1}, 1))
}
