package main

import (
	"fmt"
)

func solve(i, sum int, nums []int, dp *[][]int) bool {
	if sum == 0 {
		return true
	}
	if i == len(nums) || sum < 0 {
		return false
	}
	if (*dp)[i][sum] != -1 {
		return (*dp)[i][sum] == 1
	}
	ans := solve(i+1, sum-nums[i], nums, dp) || solve(i+1, sum, nums, dp)
	if ans {
		(*dp)[i][sum] = 1
	} else {
		(*dp)[i][sum] = 0
	}
	return ans
}

func canPartition(nums []int) bool {
	sum := 0
	for _, ele := range nums {
		sum += ele
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, sum+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return solve(0, sum, nums, &dp)
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
