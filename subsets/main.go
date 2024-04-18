package main

import "fmt"

func subsets(nums []int) [][]int {
	var ans [][]int
	ans = append(ans, []int{})
	n := len(nums)
	for i := 1; i < (1 << n); i++ {
		var temp []int
		for j := 0; j < n; j++ {
			if i&(1<<j) > 0 {
				temp = append(temp, nums[j])
			}
		}
		ans = append(ans, temp)
	}
	return ans
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
