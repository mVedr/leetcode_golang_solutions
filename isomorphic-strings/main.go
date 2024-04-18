package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	ms := make(map[byte]byte)
	mt := make(map[byte]byte)

	for i := range s {
		ca, ok1 := ms[s[i]]
		cb, ok2 := mt[t[i]]

		if ok1 != ok2 || ok1 && ca != t[i] || ok2 && cb != s[i] {
			return false
		}

		ms[s[i]] = t[i]
		mt[t[i]] = s[i]
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
}
