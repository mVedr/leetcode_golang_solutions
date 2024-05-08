package main

import (
	"fmt"
	"strconv"
)

var (
	MOD = 1_000_000_007
)

func solve(i int, currSum int, relax int, min_sum, max_sum int, num string, dp map[[3]int]int) int {
	if i >= len(num) {
		if currSum >= min_sum {
			return 1
		}
		return 0
	}
	if v, ok := dp[[3]int{i, currSum, relax}]; ok {
		return v
	}
	ans := 0

	if relax == 1 {
		for I := range 10 {
			if currSum+I <= max_sum {
				ans = (ans + solve(i+1, currSum+I, relax, min_sum, max_sum, num, dp)) % MOD
			}
		}
	} else {
		for I := range int(num[i] - '0') {
			if currSum+I <= max_sum {
				ans = (ans + solve(i+1, currSum+I, 1, min_sum, max_sum, num, dp)) % MOD
			}
		}
		nn := int(num[i] - '0')
		if currSum+nn <= max_sum {
			ans = (ans + solve(i+1, currSum+nn, relax, min_sum, max_sum, num, dp)) % MOD
		}
	}
	dp[[3]int{i, currSum, relax}] = ans % MOD
	return ans % MOD
}

func count(num1 string, num2 string, min_sum int, max_sum int) int {
	dp1, dp2 := make(map[[3]int]int), make(map[[3]int]int)
	b := solve(0, 0, 0, min_sum, max_sum, num2, dp1)
	nn, _ := strconv.Atoi(num1)
	a := solve(0, 0, 0, min_sum, max_sum, strconv.FormatInt(int64(nn-1), 10), dp2)
	return (b - a) % MOD
}

func main() {
	fmt.Println(count("1", "12", 1, 8))
	fmt.Println(count("1", "5", 1, 5))
	//fmt.Println(count(""))
}
