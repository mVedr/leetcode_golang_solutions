package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Atoi(s string) int {
	arr := strings.Fields(s)
	for _, ele := range arr {
		if val, err := strconv.ParseFloat(ele, 64); err == nil {
			return int(val)
		} else {
			return 0
		}
	}
	return 0
}

func main() {
	fmt.Println(Atoi("42"))
	fmt.Println(Atoi("   -42"))
	fmt.Println(Atoi("4193 with words"))
	fmt.Println(Atoi("words and 987"))
}
