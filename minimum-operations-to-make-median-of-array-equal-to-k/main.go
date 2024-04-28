package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solveOdd(nums []int) int64 {
	mid := len(nums) / 2
	ans := int64(0)
	if nums[mid] > 0 {
		for I := mid; I >= 0; I-- {
			if nums[I] > 0 {
				ans += int64(nums[I])
			}
		}
	} else if nums[mid] < 0 {
		for I := len(nums) - 1; I >= mid; I-- {
			if nums[I] < 0 {
				ans += int64(abs(nums[I]))
			}
		}
	}
	return ans
}

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for I := range nums {
		nums[I] -= k
	}
	return solveOdd(nums)
}

func main() {
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 4))
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 7))
	fmt.Println(minOperationsToMakeMedianK([]int{1, 2, 3, 4, 5, 6}, 4))
}
