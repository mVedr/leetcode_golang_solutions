package main

import (
	"fmt"
	"math"
	"sort"
)

func divs(n int) []int {
	ans := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			ans = append(ans, i)
			if n/i != i {
				ans = append(ans, n/i)
			}
		}
	}
	return ans
}

func minAnagramLength(s string) int {
	divisors := divs(len(s))
	//fmt.Println(divisors)
	sort.Ints(divisors)
	for _, t := range divisors {
		cnt := make([]int, 26)
		fl := true
		//fmt.Println("t: ", t)
		for i := 0; i < len(s); i += t {
			newCnt := make([]int, 26)
			//fmt.Println("cnt: ", cnt)
			for j := i; j < i+t; j++ {
				newCnt[int(s[j]-'a')]++
			}
			if i != 0 {
				for ii := 0; ii < 26; ii++ {
					if newCnt[ii] != cnt[ii] {
						fl = false
						break
					}
				}
			}
			//fmt.Println("newCnt: ", newCnt)
			cnt = newCnt
		}
		if fl {
			return t
		}
	}
	return -1
}

func main() {
	fmt.Println(minAnagramLength("abba"))
	fmt.Println(minAnagramLength("cdef"))
}
