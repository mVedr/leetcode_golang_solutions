package main

import "fmt"

func solve(I, k int, vis map[int]int, obs []int, dp map[[2]int]bool) bool {
	if v, ok := dp[[2]int{I, k}]; ok {
		return v
	}
	ans := false
	if I == 0 {
		if val, ok := vis[obs[I]+k]; ok {
			if val == len(obs)-1 {
				return true
			}
			ans = ans || solve(val, k, vis, obs, dp)
		}
	} else {
		if val, ok := vis[obs[I]+k]; ok {
			if val == len(obs)-1 {
				return true
			}
			ans = ans || solve(val, k, vis, obs, dp)
		}
		if k > 1 {

			if val, ok := vis[obs[I]+k-1]; ok {
				if val == len(obs)-1 {
					return true
				}
				ans = ans || solve(val, k-1, vis, obs, dp)
			}
		}

		if val, ok := vis[obs[I]+k+1]; ok {
			if val == len(obs)-1 {
				return true
			}
			ans = ans || solve(val, k+1, vis, obs, dp)
		}
	}

	dp[[2]int{I, k}] = ans
	return ans
}

func canCross(stones []int) bool {
	vis := make(map[int]int)
	for I, v := range stones {
		vis[v] = I
	}
	dp := make(map[[2]int]bool)
	return solve(0, 1, vis, stones, dp)
}

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
	fmt.Println(canCross([]int{0, 1, 2, 5, 6, 9, 10, 12, 13, 14, 17, 19, 20, 21, 26, 27, 28, 29, 30}))
}
