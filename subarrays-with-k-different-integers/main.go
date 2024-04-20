package main

import "fmt"

func solve(nums []int, k int) int {
	mp := map[int]int{}
	ans := 0
	cnt := 0
	I, J := 0, 0
	for J < len(nums) {
		val, ok := mp[nums[J]]
		if val == 0 || !ok {
			cnt++
		}
		mp[nums[J]]++
		for cnt > k {
			mp[nums[I]]--
			if mp[nums[I]] == 0 {
				cnt--
			}
			I++
		}
		ans += (J - I + 1)
		J++
	}
	return ans
}
func subarraysWithKDistinct(nums []int, k int) int {
	return solve(nums, k) - solve(nums, k-1)
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
	fmt.Println(subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))
}
