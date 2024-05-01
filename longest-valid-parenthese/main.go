package main

import "fmt"

type Node struct {
	ind int
	typ int
}

func longestValidParentheses(s string) int {
	ans := 0
	st := []Node{}
	st = append(st, Node{ind: -1, typ: 1})
	for I, ch := range s {
		if ch == '(' {
			st = append(st, Node{ind: I, typ: 0})
		} else {
			tp := st[len(st)-1]
			if tp.typ == 0 {
				st = st[:len(st)-1]
				ans = max(ans, I-st[len(st)-1].ind)
			} else {
				st = append(st, Node{typ: 1, ind: I})
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses(""))
}
