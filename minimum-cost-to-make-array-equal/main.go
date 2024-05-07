package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func cos(mid int, nums, cost []int) int64 {
	c := int64(0)
	for I := range nums {
		c += int64(abs(nums[I]-mid) * cost[I])
	}
	return c
}

func minCost(nums []int, cost []int) int64 {
	lo := 0
	hi := 1_000_000
	for lo < hi {
		mid := lo + (hi-lo)/2
		mm, ll, uu := cos(mid, nums, cost), cos(mid-1, nums, cost), cos(mid+1, nums, cost)
		if mm <= ll && mm <= uu {
			return mm
		} else if mm < uu {
			hi = mid
		} else if mm > uu {
			lo = mid + 1
		}
	}
	return int64(-1)
}

func main() {
	fmt.Println(minCost([]int{1, 3, 5, 2}, []int{2, 3, 1, 4}))
	fmt.Println(minCost([]int{2, 2, 2, 2, 2}, []int{4, 2, 8, 1, 3}))

}
