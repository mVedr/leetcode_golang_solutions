package main

import (
	"fmt"
	"sort"
)

func maxPalindromesAfterOperations(words []string) int {
	cnt, pp, ans := [26]int{}, 0, 0
	sizes := []int{}
	for _, w := range words {
		for _, ch := range w {
			if (cnt[int(ch-'a')]+1)%2 == 0 {
				pp++
			}
			cnt[int(ch-'a')]++
		}
		sizes = append(sizes, len(w))
	}
	sort.Ints(sizes)
	for _, sz := range sizes {
		pp -= sz / 2
		if pp >= 0 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(maxPalindromesAfterOperations([]string{"abbb", "ba", "aa"}))
	fmt.Println(maxPalindromesAfterOperations([]string{"abc", "ab"}))
	fmt.Println(maxPalindromesAfterOperations([]string{"cd", "ef", "a"}))
}
