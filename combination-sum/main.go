package main

import "fmt"

func solve(I int, t int, arr []int, ans *[][]int, temp []int) {
	if t == 0 {
		cpy := make([]int, len(temp))
		copy(cpy, temp)
		*ans = append(*ans, cpy)
		return
	}
	if I >= len(arr) {
		return
	}
	if t-arr[I] >= 0 {
		temp = append(temp, arr[I])
		solve(I, t-arr[I], arr, ans, temp)
		temp = temp[:len(temp)-1]
	}
	solve(I+1, t, arr, ans, temp)
}

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	solve(0, target, candidates, &ans, []int{})
	return ans
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}
