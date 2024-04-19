package main

import "fmt"

func rearrangeArray(nums []int) []int {
	a, b, n := 0, 1, len(nums)
	ans := make([]int, n)
	for _, ele := range nums {
		if ele > 0 {
			ans[a] = ele
			a += 2
		}
	}
	for _, ele := range nums {
		if ele < 0 {
			ans[b] = ele
			b += 2
		}
	}
	return ans
}

func main() {
	fmt.Println(rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
	fmt.Println(rearrangeArray([]int{-1, 1}))
}
