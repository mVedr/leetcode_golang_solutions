package main

import (
	"fmt"
	"math"
	"sort"
)

func ff(freq []int, i, target int) int {
	l := i
	r := len(freq)
	for l < r {
		mid := (l + r) / 2
		if freq[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}

func solve(freq []int, k int) int {
	ans := math.MaxInt32
	pre := make([]int, len(freq)+1)
	sum := 0
	for i := 0; i < len(freq); i++ {
		sum += freq[i]
		pre[i+1] = sum
	}

	for i := 0; i < len(freq); i++ {
		curr := 0
		target := freq[i] + k
		ffalse := ff(freq, i, target)
		cnt := len(freq) - ffalse
		rightSum := pre[len(freq)] - pre[ffalse]
		curr += (rightSum - cnt*target)
		curr += pre[i]
		ans = min(ans, curr)
	}

	return ans
}

func minimumDeletions(word string, k int) int {
	freq := make([]int, 26)
	for i := range word {
		c := word[i] - 'a'
		freq[c]++
	}

	sort.Ints(freq)
	ans := solve(freq, k)
	return ans
}

func main() {
	fmt.Println(minimumDeletions("aabcaba", 0))
	fmt.Println(minimumDeletions("dabdcbdcdcd", 2))
	fmt.Println(minimumDeletions("aaabaaa", 2))
}
