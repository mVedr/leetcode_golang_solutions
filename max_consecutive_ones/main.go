package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	ans := 0
	i, j := 0, 0
	for j < n {
		if arr[j] == 1 {
			j++
			ans = max(ans, j-i)
		} else {
			i = j + 1
			j++
		}
	}
	fmt.Print(ans)
}
