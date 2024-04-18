package main

import "fmt"

func solveZ(nums []int) {}

func solve(nums []int, k int) int {
	mpp := make(map[int]int)
	mpp[0] = 1
	ans := 0
	sum := 0
	for _, ele := range nums {
		sum += ele
		if cnt, ok := mpp[sum-k]; ok {
			ans += cnt
		}
		mpp[sum]++
	}

	return ans
}

func main() {
	var n, goal int
	fmt.Scan(&n, &goal)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	fmt.Println(solve(nums, goal))
}
