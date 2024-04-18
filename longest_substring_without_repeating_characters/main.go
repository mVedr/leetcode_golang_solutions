package main

import "fmt"

func solve(str string) int {
	v := make([]int, 1024)
	for in := range v {
		v[in] = -1
	}
	ans := 0
	currIndex := -1
	for i := 0; i < len(str); i++ {
		if v[str[i]] > currIndex {
			currIndex = v[str[i]]
		}
		v[str[i]] = i
		ans = max(ans, i-currIndex)
	}
	return ans
}

func main() {
	fmt.Println(solve("abcabcbb"))
	fmt.Println(solve("bbbbb"))
	fmt.Println(solve("pwwkew"))
}
