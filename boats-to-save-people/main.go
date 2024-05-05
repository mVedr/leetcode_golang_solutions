package main

import "sort"

func ok(mid int, people []int, limit int) bool {
	currSum := 0
	cnt := 1
	for _, p := range people {
		if currSum+p > limit {
			currSum = p
			cnt++
		}
	}
	return cnt <= mid
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	i, j := 0, len(people)-1
	cnt := 0
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
			j--
		} else {
			j--
		}
		cnt++
	}
	return cnt
}

func main() {

}
