package main

import "fmt"

func uniqueLetterString(s string) int {
	l, r := make([][26]int, len(s)), make([][26]int, len(s))
	ans := 0
	for I := 0; I < 26; I++ {
		l[0][I] = -1
		r[len(s)-1][I] = len(s)
	}
	for I := 0; I < len(s); I++ {
		l[I][byte(s[I])-'A'] = I
	}

	for I := len(s) - 1; I >= 0; I-- {
		r[I][byte(s[I])-'A'] = I
	}

	for I := 0; I < len(s); I++ {
		if I == 0 {
			ans += (r[I+1][byte(s[I])-'A'] - I)
		} else if I == len(s)-1 {
			ans += (I - l[I-1][byte(s[I])-'A'])
		} else {
			ans += ((r[I+1][byte(s[I])-'A'] - I) * (I - l[I-1][byte(s[I])-'A']))
		}
	}

	return ans
}

func main() {
	fmt.Println(uniqueLetterString("ABC"))
	fmt.Println(uniqueLetterString("ABA"))
	fmt.Println(uniqueLetterString("LEETCODE"))
}
