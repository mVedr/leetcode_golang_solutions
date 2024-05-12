package main

import "fmt"

const (
	MOD = 1_000_000_007
)

func f(msk int, prev int, n int, nums []int, dp map[[2]int]int) int {
	if msk == (1<<n)-1 {
		return 1
	}
	if v, ok := dp[[2]int{msk, prev}]; ok {
		return v
	}
	ans := 0
	for I := 0; I < n; I++ {
		if 1<<I&msk == 0 {
			if prev == -1 {
				ans = (ans + f(msk|1<<I, I, n, nums, dp)) % MOD
			} else if nums[I]%nums[prev] == 0 || nums[prev]%nums[I] == 0 {
				ans = (ans + f(msk|1<<I, I, n, nums, dp)) % MOD
			}
		}
	}
	dp[[2]int{msk, prev}] = ans % MOD
	return ans % MOD
}

func specialPerm(nums []int) int {
	dp := make(map[[2]int]int)
	return f(0, -1, len(nums), nums, dp)
}

func main() {
	fmt.Println(specialPerm([]int{2, 3, 6}))
	fmt.Println(specialPerm([]int{1, 4, 3}))
}
