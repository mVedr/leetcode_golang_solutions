package main

import "fmt"

func minSum(nums1 []int, nums2 []int) int64 {

	c1, c2 := 0, 0
	s1, s2 := int64(0), int64(0)
	for _, N := range nums1 {
		if N == 0 {
			c1++
		}
		s1 += int64(N)
	}
	for _, N := range nums2 {
		if N == 0 {
			c2++
		}
		s2 += int64(N)
	}
	if c1 == 0 && s1 < s2+int64(c2) {
		return -1
	}
	if c2 == 0 && s2 < s1+int64(c1) {
		return -1
	}
	return max(s2+int64(c2), s1+int64(c1))

}
func main() {
	fmt.Println(minSum([]int{3, 2, 0, 1, 0}, []int{6, 5, 0}))
	fmt.Println(minSum([]int{2, 0, 2, 0}, []int{1, 4}))
}
