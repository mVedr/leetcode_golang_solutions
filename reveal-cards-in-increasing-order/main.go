package main

import (
	"fmt"
	"slices"
)

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	slices.Sort(deck)
	res := make([]int, n)
	q := make([]int, n)

	for i := range q {
		q[i] = i
	}

	for _, val := range deck {
		idx := q[0]
		q = q[1:]
		res[idx] = val

		if len(q) > 0 {
			q = append(q, q[0])
			q = q[1:]
		}
	}
	return res
}

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17, 13, 11, 2, 3, 5, 7}))
	fmt.Println(deckRevealedIncreasing([]int{1, 1000}))
}
