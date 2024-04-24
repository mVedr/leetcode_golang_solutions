package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	st := []int{}
	mpp := map[int]int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(st) > 0 && st[len(st)-1] <= nums2[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			mpp[nums2[i]] = -1
		} else {
			mpp[nums2[i]] = st[len(st)-1]
		}
		st = append(st, nums2[i])
	}
	ans := make([]int, len(nums1))
	for i, ele := range nums1 {
		ans[i] = mpp[ele]
	}
	return ans
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
