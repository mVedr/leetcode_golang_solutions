package main

import (
	"fmt"
	"reflect"
)

func solve(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	b += b
	for i := range len(b) / 2 {
		ok := reflect.DeepEqual(a, b[i:i+len(a)])
		if ok {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(solve("abcde", "cdeab"))
	fmt.Println(solve("abcde", "abced"))
}
