package main

import "fmt"

func solve(s string) string {
	ans := []byte{}
	stk := []byte{}
	for _, ch := range s {
		if ch == '(' {
			stk = append(stk, byte(ch))
			if len(stk) > 1 {
				ans = append(ans, byte(ch))
			}
		} else {
			stk = stk[:len(stk)-1]
			if len(stk) > 0 {
				ans = append(ans, byte(ch))
			}
		}
	}
	return string(ans)
}

func main() {
	fmt.Println(solve("(()())(())"))
	fmt.Println(solve("(()())(())(()(()))"))
}
