package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	mpp := make(map[int]int)
	mpp[0] = 1
	ans := 0
	sum := 0
	for _, ele := range nums {
		sum += ele
		if cnt, ok := mpp[sum-k]; ok {
			ans += cnt
		}
		mpp[sum]++
	}

	fmt.Println(ans)

}
