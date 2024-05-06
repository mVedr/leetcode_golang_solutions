package main

import "fmt"

func minimumOperationsToMakeKPeriodic(word string, k int) int {
	if k == len(word) {
		return 0
	}
	mp := make(map[string]int)
	I := 0
	for I < len(word) {
		fmt.Println(word[I : I+k])
		mp[word[I:I+k]]++
		I += k
	}
	fmt.Println(mp)
	mx := 0
	str := ""
	for k, v := range mp {
		fmt.Println(k, v)
		if v > mx {
			mx = v
			str = k
		}
	}
	ans := 0
	for k, v := range mp {
		if k != str {
			ans += v
		}
	}

	return ans
}

func main() {
	fmt.Println(minimumOperationsToMakeKPeriodic("xu", 1))
}
