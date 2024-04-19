package main

import (
	"fmt"
	"math"
)

func ok(sp, h int, pile []int) bool {
	tm := 0
	for _, ele := range pile {
		num := float64(ele) / float64(sp)
		tm += int(math.Ceil(num))
		if tm > h {
			return false
		}
	}
	return tm <= h
}

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, math.MaxInt32
	ans := hi
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if ok(mid, h, piles) {
			hi = mid - 1
			ans = min(ans, mid)
		} else {
			lo = mid + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}
