package main

import (
	"fmt"
	"reflect"
	"sort"
)

func solve(s string, t string) bool {
	st := []rune(s)
	tt := []rune(t)
	sort.Slice(st, func(i int, j int) bool { return st[i] < st[j] })
	sort.Slice(tt, func(i int, j int) bool { return tt[i] < tt[j] })
	return reflect.DeepEqual(st, tt)
}

func main() {
	fmt.Println(solve("nagaram", "anagram"))
	fmt.Println(solve("car", "rat"))
}
