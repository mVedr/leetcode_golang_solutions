package main

import (
	"fmt"
	"math"
)

func scoreOfString(s string) int {
	ans := 0
	for i := range len(s) - 1 {
		ans += int(math.Abs(float64(int8(byte(s[i])) - int8(byte(s[i+1])))))
	}
	return ans
}

func main() {
	fmt.Println(scoreOfString("hello"))
	fmt.Println(scoreOfString("zaz"))
}
