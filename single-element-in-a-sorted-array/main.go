package main

import "fmt"

func singleNonDuplicate(nums []int) int {

	lo, hi, n := 0, len(nums)-1, len(nums)
	if n == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[n-2] != nums[n-1] {
		return nums[n-1]
	}
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if mid%2 == 0 {
			if mid+1 < n && nums[mid+1] == nums[mid] {
				hi = mid + 2
			} else if mid >= 1 && nums[mid-1] == nums[mid] {
				lo = mid - 2
			} else {
				return nums[mid]
			}
		} else {
			if mid+1 < n && nums[mid+1] == nums[mid] {
				hi = mid - 1
			} else if mid >= 1 && nums[mid-1] == nums[mid] {
				lo = mid + 1
			} else {
				return nums[mid]
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3}))
}
