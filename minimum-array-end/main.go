package main

import "fmt"

func minEnd(n int, x int) int64 {
	ans := int64(0)
	z := 0
	mp := make(map[int]int)
	for i := int64(0); i < 64; i++ {
		if (1<<i)&x > 0 {
			ans = ans | (1 << i)
		} else {
			mp[z] = int(i)
			z++
		}
	}
	lastEl := n - 1
	for i := int64(0); i <= 63; i++ {
		if (1<<i)&lastEl > 0 {
			ans = ans | (1 << mp[int(i)])
		}
	}
	return ans
}

func main() {
	fmt.Println(minEnd(3, 4))
	fmt.Println(minEnd(2, 7))
}
