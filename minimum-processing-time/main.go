package main

import "sort"

func minProcessingTime(processorTime []int, tasks []int) int {
	sort.Ints(tasks)
	sort.Slice(processorTime, func(i, j int) bool {
		return processorTime[i] > processorTime[j]
	})
	ans := -1
	for I := 0; I < len(tasks); {
		ans = max(ans, processorTime[I/4]+tasks[I+3])
		I += 4
	}
	return ans
}

func main() {

}
