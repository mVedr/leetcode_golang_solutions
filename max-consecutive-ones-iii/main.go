package main

import "fmt"

func solve(nums []int, k int) int {
	n, i, j := len(nums), 0, 0
	ans := 0
	cnt := 0
	for j < n {
		if nums[j] == 0 {
			cnt++
		}
		for cnt > k {
			if nums[i] == 0 {
				cnt--
			}
			i++
		}
		ans = max(ans, j-i+1)
		j++
	}

	return ans
}

func main() {
	fmt.Println(solve([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	fmt.Println(solve([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
