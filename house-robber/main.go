package main

func solve(i int, nums []int, dp *[]int) int {
	if i >= len(nums) {
		return 0
	}
	if (*dp)[i] != -1 {
		return (*dp)[i]
	}
	ans := 0
	ans = max(ans, nums[i]+solve(i+2, nums, dp))
	ans = max(ans, solve(i+1, nums, dp))
	(*dp)[i] = ans
	return ans
}

func rob(nums []int) int {
	dp := make([]int, 105)
	for I := range len(dp) {
		dp[I] = -1
	}
	return solve(0, nums, &dp)
}

func main() {

}
