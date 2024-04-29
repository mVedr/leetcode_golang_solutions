package main

import "fmt"

func minOperations(nums []int, k int) int {
	num := [32]int{}
	for _, N := range nums {
		for I := 0; I < 32; I++ {
			if (1<<I)&N > 0 {
				num[I]++
			}
		}
	}
	ans := 0
	for I := 0; I < 32; I++ {
		if (1<<I)&k > 0 {
			if num[I]%2 == 0 {
				ans++
			}
		} else {
			if num[I]%2 == 1 {
				ans++
			}
		}
	}
	//fmt.Println(num)
	return ans
}

func main() {
	fmt.Println(minOperations([]int{2, 1, 3, 4}, 1))
	fmt.Println(minOperations([]int{2, 0, 2, 0}, 0))
}
