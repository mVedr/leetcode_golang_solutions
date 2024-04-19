package main

import (
	"fmt"
	"math"
	"slices"
)

func allNeg(nums []int) bool {
	for _, ele := range nums {
		if ele >= 0 {
			return false
		}
	}
	return true
}

func maxSubArray(nums []int) int {
	ans := math.MinInt32
	sum := 0
	mx := slices.MaxFunc(nums, func(i, j int) int {
		return i - j
	})
	if allNeg(nums) {
		return mx
	}
	for _, ele := range nums {
		if sum+ele < 0 {
			sum = 0
		} else {
			sum += ele
		}
		ans = max(ans, sum)
	}
	return ans
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{-1}))
}
