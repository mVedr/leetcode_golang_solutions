package main

import (
	"sort"
)

func canSeePersonsCount(heights []int) []int {
	n := len(heights)
	st := []int{heights[n-1]}
	ans := make([]int, n)

	for i := n - 2; i >= 0; i-- {
		idx := sort.Search(len(st),
			func(j int) bool { return st[j] < heights[i] })
		if idx > 0 {
			idx--
		}
		ele := len(st) - idx
		ans[i] = ele
		for len(st) > 0 && st[len(st)-1] <= heights[i] {
			st = st[:len(st)-1]
		}
		st = append(st, heights[i])
	}
	return ans
}
func main() {

}
