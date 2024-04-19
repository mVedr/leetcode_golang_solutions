package main

import "fmt"

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	ans := len(nums)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			ans = min(ans, mid)
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
