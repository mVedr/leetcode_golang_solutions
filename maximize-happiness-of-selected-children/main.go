package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	for I := range happiness {
		happiness[I] = happiness[I] - k
	}
	sub := k
	ans := int64(0)
	for I := len(happiness) - 1; I >= 0; I-- {
		if k > 0 {
			ans += max(int64(happiness[I]+sub), 0)
			k--
			sub--
		}
	}
	return ans
}

func main() {
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))
	fmt.Println(maximumHappinessSum([]int{1, 1, 1, 1}, 2))
	fmt.Println(maximumHappinessSum([]int{2, 3, 4, 5}, 1))
}
