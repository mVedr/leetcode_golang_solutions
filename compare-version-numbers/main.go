package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")
	nums1, nums2 := []int{}, []int{}
	for _, v := range arr1 {
		n, _ := strconv.Atoi(v)
		nums1 = append(nums1, n)
	}
	for _, v := range arr2 {
		n, _ := strconv.Atoi(v)
		nums2 = append(nums2, n)
	}
	if len(arr1) < len(arr2) {
		for J := len(arr1); J < len(arr2); J++ {
			nums1 = append(nums1, 0)
		}
	}
	if len(arr2) < len(arr1) {
		for J := len(arr2); J < len(arr1); J++ {
			nums2 = append(nums2, 0)
		}
	}
	for I := range nums1 {
		if nums1[I] < nums2[I] {
			return -1
		}
		if nums2[I] < nums1[I] {
			return 1
		}
	}
	return 0
}

func main() {

}
