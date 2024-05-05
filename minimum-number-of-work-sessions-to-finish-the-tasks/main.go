package main

import "math"

func solve(mask, currTime int, tasks []int, sessionTime int,
	dp map[[2]int]int) int {
	if mask == 1<<(len(tasks))-1 {
		return 0
	}

	if v, ok := dp[[2]int{mask, currTime}]; ok {
		return v
	}

	ans := math.MaxInt32
	for I := range tasks {
		if mask&(1<<I) > 0 {
			//already
		} else {
			if tasks[I]+currTime > sessionTime {
				ans = min(ans, 1+solve(mask|(1<<I), tasks[I], tasks, sessionTime, dp))
			} else {
				ans = min(ans, solve(mask|(1<<I), tasks[I]+currTime, tasks, sessionTime, dp))
				ans = min(ans, 1+solve(mask|(1<<I), tasks[I], tasks, sessionTime, dp))
			}
		}
	}
	dp[[2]int{mask, currTime}] = ans
	return ans
}

func minSessions(tasks []int, sessionTime int) int {
	dp := make(map[[2]int]int)
	return solve(0, 0, tasks, sessionTime, dp) + 1
}

func main() {

}
