package main

import "fmt"

func solve(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

func main() {
	fmt.Println(solve(4))
	fmt.Println(solve(14))
	fmt.Println(solve(-4))
}
