package main

import (
	"fmt"
	"math"
)

func solve(x, y, z, prev int, dp map[[4]int]int) int {
	if x == 0 && y == 0 && z == 0 {
		return 0
	}
	if x < 0 || y < 0 || z < 0 {
		return math.MinInt16
	}
	if val, ok := dp[[4]int{x, y, z, prev}]; ok {
		return val
	}

	ans := 0
	if prev == -1 {
		ans = max(ans, 2+solve(x-1, y, z, 0, dp))
		ans = max(ans, 2+solve(x, y-1, z, 1, dp))
		ans = max(ans, 2+solve(x, y, z-1, 2, dp))
	} else {
		if prev == 0 { //AA
			ans = max(ans, 2+solve(x, y-1, z, 1, dp))
		} else if prev == 1 { //BB
			ans = max(ans, 2+solve(x-1, y, z, 0, dp))
			ans = max(ans, 2+solve(x, y, z-1, 2, dp))
		} else { //AB
			ans = max(ans, 2+solve(x-1, y, z, 0, dp))
			ans = max(ans, 2+solve(x, y, z-1, 2, dp))
		}
	}
	dp[[4]int{x, y, z, prev}] = ans
	return ans
}

func longestString(x int, y int, z int) int {
	dp := make(map[[4]int]int)
	return solve(x, y, z, -1, dp)
}

func main() {
	fmt.Println(longestString(2, 5, 1))
	fmt.Println(longestString(3, 2, 2))
}
