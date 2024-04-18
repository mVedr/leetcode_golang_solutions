package main

import "fmt"

func atMost(nums []int, k int) int {
	ans := 0
	cnt := 0
	i := 0
	j := 0
	for j < len(nums) {
		if nums[j]%2 == 1 {
			cnt++
		}
		for cnt > k {
			if nums[i]%2 == 1 {
				cnt--
			}
			i++
		}
		ans += (j - i + 1)
		j++
	}
	return ans
}

func solve(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
}

func main() {
	fmt.Println(solve([]int{1, 1, 2, 1, 1}, 3))
	fmt.Println(solve([]int{2, 4, 6}, 1))
	fmt.Println(solve([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
}
