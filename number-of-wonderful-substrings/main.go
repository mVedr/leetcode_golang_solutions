package main

import (
	"fmt"
)

func wonderfulSubstrings(word string) int64 {
	cnt := [1025]int64{}
	cnt[0] = 1
	mask := 0
	ans := int64(0)
	for _, c := range word {
		ind := int(c - 'a')
		mask = mask ^ (1 << ind)
		ans += cnt[mask]
		for i := range 10 {
			temp := mask ^ (1 << i)
			ans += cnt[temp]
		}
		cnt[mask]++
	}

	return ans
}
func main() {
	fmt.Println(wonderfulSubstrings("aba"))
	fmt.Println(wonderfulSubstrings("aabb"))
	fmt.Println(wonderfulSubstrings("he"))
}
