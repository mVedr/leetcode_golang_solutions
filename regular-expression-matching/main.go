package main

import "fmt"

func ok(ch byte) bool {
	return ch >= 'a' && ch <= 'z'
}

func solve(I, J int, s, p string, dp map[[2]int]bool) bool {
	if I == -1 && J == -1 {
		return true
	}
	if J == -1 {
		return false
	}
	if I == -1 {
		for K := J; K >= 0; K-- {
			if p[K] != '*' {
				return false
			}
			K--
		}
		return true
	}

	if val, ok := dp[[2]int{I, J}]; ok {
		return val
	}

	ans := false
	if ok(p[J]) {
		if s[I] != p[J] {
			ans = ans || false
		} else {
			ans = ans || solve(I-1, J-1, s, p, dp)
		}
	} else if p[J] == '.' {
		ans = ans || solve(I-1, J-1, s, p, dp)
	} else if p[J] == '*' {
		ans = ans || solve(I, J-2, s, p, dp)

		if s[I] == p[J-1] || p[J-1] == '.' {
			ans = ans || solve(I-1, J, s, p, dp)
		}
	} else {
		ans = ans || false
	}
	dp[[2]int{I, J}] = ans
	return ans
}

func isMatch(s string, p string) bool {
	dp := make(map[[2]int]bool)
	return solve(len(s)-1, len(p)-1, s, p, dp)
}

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".*"))
}
