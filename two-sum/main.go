package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	var target int
	fmt.Scan(&target)
	mpp := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	for ind, num := range nums {
		if pos, ok := mpp[target-num]; ok {
			fmt.Println(pos, " ", ind)
		}
		mpp[num] = ind
	}
}
