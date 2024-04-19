package main

import (
	"fmt"
	"sort"
)

func minRectanglesToCoverPoints(points [][]int, w int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})
	ans := 0
	curr := 0
	for i := 1; i < len(points); i++ {
		if points[i][0]-points[curr][0] > w {
			ans++
			curr = i
		}
	}
	if curr != len(points) {
		ans++
	}
	return ans
}

func main() {
	fmt.Println(minRectanglesToCoverPoints([][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}, 1))
	fmt.Println(minRectanglesToCoverPoints([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}, 2))
	fmt.Println(minRectanglesToCoverPoints([][]int{{2, 3}, {1, 2}}, 0))
}
