package main

import "sort"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func ok(a, b int) bool {
	return a < 0 && b > 0
}

func findMaxK(nums []int) int {
	sort.Ints(nums)
	I, J := 0, len(nums)-1
	for I < J {
		if !ok(nums[I], nums[J]) {
			break
		}
		if abs(nums[I]) > nums[J] {
			I++
		} else if abs(nums[I]) < nums[J] {
			J--
		} else {
			return nums[J]
		}
	}
	return -1
}

func main() {

}
