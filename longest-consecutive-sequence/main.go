package main

import (
	"fmt"
	"slices"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	slices.SortFunc(nums, func(a int, b int) int { return a - b })

	i := 0
	ans := 0
	for i < n {
		j := i + 1
		for (j < n) && (nums[j]-nums[j-1] == 1) {
			ans = max(ans, (j - i + 1))
			j++
		}
		i = j
	}
	fmt.Printf("%d\n", ans)
}
