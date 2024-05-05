package main

import "sort"

func countWays(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for I, v := range nums {
		take := I + 1
		if take > v {
			ans++
		}

	}
	return ans
}

func main() {

}
