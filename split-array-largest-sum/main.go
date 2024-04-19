package main

import (
	"fmt"
	"math"
)

func ok(sum, k int, nums []int) bool {
	count := 1
	curSum := 0
	for _, num := range nums {
		if num > sum {
			return false
		}
		curSum += num
		if curSum > sum {
			count++
			curSum = num
		}
	}
	return count <= k
}

func splitArray(nums []int, k int) int {
	lo, hi := 0, math.MaxInt32
	ans := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if ok(mid, k, nums) {
			ans = min(ans, mid)
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println(splitArray([]int{1, 2, 3, 4, 5}, 2))
}
