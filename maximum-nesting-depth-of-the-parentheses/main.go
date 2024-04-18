package main

import "fmt"

func solve(s string) int {
	ans := 0
	cnt := 0
	for _, ch := range s {
		if ch == '(' {
			cnt++
			ans = max(ans, cnt)
		} else if ch == ')' {
			cnt--
		}
	}
	return ans
}

func main() {
	fmt.Println(solve("(1+(2*3)+((8)/4))+1"))
	fmt.Println(solve("(1)+((2))+(((3)))"))
}
