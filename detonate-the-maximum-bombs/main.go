package main

import (
	"fmt"
	"math"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func solve(curr int, bombs [][]int) int {
	vis := make(map[int]bool)
	q := []int{}
	q = append(q, curr)
	vis[curr] = true
	ans := 0
	for len(q) > 0 {
		n := len(q)
		for range n {
			tp := q[0]
			q = q[1:]
			for I, arr := range bombs {
				if !vis[I] {
					r := float64(bombs[tp][2])
					if distance(arr[0], arr[1], bombs[tp][0], bombs[tp][1]) <= r {
						q = append(q, I)
						vis[I] = true
						ans++
					}
				}
			}
		}
	}
	return ans + 1
}

func maximumDetonation(bombs [][]int) int {
	ans := 0
	for arr := range bombs {
		ans = max(ans, solve(arr, bombs))
	}
	return ans
}

func main() {
	fmt.Println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
	fmt.Println(maximumDetonation([][]int{{1, 1, 5}, {10, 10, 5}}))
	fmt.Println(maximumDetonation([][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}))
}
