package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	ans := 0
	curr := 0
	for _, ele := range g {
		if curr >= len(s) {
			break
		}
		if ele <= s[curr] {
			ans++
			curr++
		} else {
			for curr < len(s) && ele > s[curr] {
				curr++
			}
			if curr >= len(s) {
				break
			} else {
				ans++
				curr++
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
