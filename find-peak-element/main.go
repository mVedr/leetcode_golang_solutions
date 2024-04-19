package main

import "fmt"

func findPeakElement(nums []int) int {
	ans := -1
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			lo = mid + 1
		} else if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			ans = mid
			hi = mid - 1
		} else {
			return mid
		}
	}
	if ans == -1 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
