package main

import "fmt"

func f(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	ans := len(nums)

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			ans = min(ans, mid)
			hi = mid - 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if ans == len(nums) {
		return -1
	}
	return ans
}

func l(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	ans := -1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			ans = max(ans, mid)
			lo = mid + 1
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func searchRange(nums []int, target int) []int {
	return []int{f(nums, target), l(nums, target)}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{}, 0))
}
