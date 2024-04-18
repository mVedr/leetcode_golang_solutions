package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := n * (n + 1) / 2

	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		sum -= temp
	}

	fmt.Print(sum)
}
