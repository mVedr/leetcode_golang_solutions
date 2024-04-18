package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	cnt := 0
	mj := -1_9 - 1
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	for _, ele := range nums {
		if mj == -1_9-1 {
			mj = ele
			cnt = 1
		} else {
			if mj != ele {
				cnt--
				if cnt == 0 {
					mj = ele
					cnt++
				}
			} else {
				cnt++
			}
		}
	}
	fmt.Printf("\n %d \n", mj)
}
