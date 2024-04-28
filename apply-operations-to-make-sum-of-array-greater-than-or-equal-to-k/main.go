package main

import (
	"fmt"
	"math"
)

func minOperations(k int) int {
	ans := math.MaxInt64
	for X := range 1_000 {
		for Y := range 1_000 {
			if (X+1)*(Y-1) >= k {
				ans = min(ans, X+Y-2)
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(minOperations(11))
	fmt.Println(minOperations(1))
	fmt.Println(minOperations(3))
}
