package main

import "fmt"

func smallestString(s string) string {
	buf := []byte(s)
	I := 0

	for I < len(s) && s[I] == 'a' {
		I++
	}

	if I == len(s) {
		buf[len(s)-1] = 'z'
		return string(buf)
	}

	for I < len(s) && s[I] != 'a' {
		buf[I] = buf[I] - 1
		I++
	}

	return string(buf)
}

func main() {
	fmt.Println(smallestString("cbabc"))
	fmt.Println(smallestString("acbbc"))
	fmt.Println(smallestString("leetcode"))
	fmt.Println(smallestString("aaaa"))

}
