package main

import "fmt"

func tribonacci(n int) int {
	p1, p2, p3 := 0, 1, 1
	if n == 0 {
		return p1
	} else if n == 1 {
		return p2
	} else if n == 2 {
		return p3
	} else {
		ans := p3
		for I := 3; I <= n; I++ {
			ans = p1 + p2 + p3
			p1 = p2
			p2 = p3
			p3 = ans
		}
		return ans
	}
}

func main() {
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(25))
}
