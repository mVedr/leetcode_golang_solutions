package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	ans := 0
	for _, x := range nums {
		ans = ans ^ x
	}
	fmt.Println(ans)
}
