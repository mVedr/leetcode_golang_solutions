package main

import (
	"math"
	"sort"
)

func ok(nums1, nums2 []int, k int) bool {
	cc := 0
	j := 0
	for i := 0; i < len(nums1) && j < len(nums2); i++ {
		if nums1[i]+k != nums2[j] {
			cc++
		} else {
			j++
		}
	}
	if cc > 2 {
		return false
	}
	return true
}

func minimumAddedInteger(nums1, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	ans := math.MaxInt32
	for i := 0; i < len(nums1); i++ {
		diff := nums2[0] - nums1[i]
		if ok(nums1, nums2, diff) {
			if diff < ans {
				ans = diff
			}
		}
	}
	return ans
}

func main() {

}
