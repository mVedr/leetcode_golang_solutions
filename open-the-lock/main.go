package main

import (
	"fmt"
)

type Node struct {
	comb  string
	steps int
}

func openLock(deadends []string, target string) int {
	nn := make(map[string]bool)
	for _, d := range deadends {
		nn[d] = true
	}
	if target == "0000" {
		return 0
	}
	if nn["0000"] || nn[target] {
		return -1
	}
	q := []Node{{"0000", 0}}
	vis := map[string]bool{"0000": true}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		for i := 0; i < 4; i++ {
			nxt := []rune(node.comb)
			nxt[i] = (nxt[i]-'0'+1)%10 + '0'
			strn := string(nxt)

			if !vis[strn] && !nn[strn] {
				if strn == target {
					return node.steps + 1
				}
				vis[strn] = true
				q = append(q, Node{strn, node.steps + 1})
			}

			prev := []rune(node.comb)
			prev[i] = (prev[i]-'0'+9)%10 + '0'
			strp := string(prev)

			if !vis[strp] && !nn[strp] {
				if strp == target {
					return node.steps + 1
				}
				vis[strp] = true
				q = append(q, Node{strp, node.steps + 1})
			}
		}
	}
	return -1
}

func main() {
	tc := []struct {
		deadends []string
		target   string
	}{
		{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202"},
		{[]string{"8888"}, "0009"},
		{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"},
	}

	for _, tc := range tc {
		fmt.Println(openLock(tc.deadends, tc.target))
	}
}
