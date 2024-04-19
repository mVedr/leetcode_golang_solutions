package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	mn := math.MaxInt32
	ans := 0
	for _, ele := range prices {
		ans = max(ans, ele-mn)
		mn = min(ele, mn)

	}
	return ans
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
