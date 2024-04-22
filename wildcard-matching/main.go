package main

import "fmt"

func solve(i, j int, s, p string, dp map[[2]int]bool) bool {
	if val, ok := dp[[2]int{i, j}]; ok {
		return val
	}

	if i < 0 && j < 0 {
		return true
	}
	if i < 0 {
		for k := j; k >= 0; k-- {
			if p[k] != '*' {
				return false
			}
		}
		return true
	}
	if j < 0 {
		return false
	}

	ans := false
	if p[j] == '*' {
		for k := i; k >= -1; k-- {
			ans = ans || solve(k, j-1, s, p, dp)
		}
	} else if p[j] == '?' || s[i] == p[j] {
		ans = ans || solve(i-1, j-1, s, p, dp)
	}

	dp[[2]int{i, j}] = ans
	return ans
}

func isMatch(s string, p string) bool {
	dp := make(map[[2]int]bool)
	return solve(len(s)-1, len(p)-1, s, p, dp)
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("cb", "?a"))
	fmt.Println(isMatch("adceb", "*a*b"))
}
