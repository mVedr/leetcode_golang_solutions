package main

import (
	"bytes"
	"fmt"
)

func getSmallestString(s string, k int) string {
	buff := new(bytes.Buffer)
	for _, c := range s {
		if k > 0 {
			mv1, mv2 := int(c-'a'), (int('a'-c)+26)%26
			if min(mv1, mv2) <= k {
				k -= min(mv1, mv2)
				buff.WriteRune('a')
			} else {
				c1, c2 := rune((int(c-'a')-k)%26+'a'), rune((int(c-'a')+k)%26+'a')
				k = 0
				buff.WriteRune(min(c1, c2))
			}

		} else {
			buff.WriteRune(c)
		}
	}
	return buff.String()
}

func main() {
	fmt.Println(getSmallestString("zbbz", 3))
	fmt.Println(getSmallestString("xaxcd", 4))
	fmt.Println(getSmallestString("lol", 0))
}
