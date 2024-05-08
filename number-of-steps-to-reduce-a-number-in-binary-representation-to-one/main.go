package main

import (
	"fmt"
)

func numSteps(s string) int {
	ans := 0
	bs := []byte(s)
	for !(len(bs) == 1 && bs[0] == '1') {
		if bs[len(bs)-1] == '0' {
			bs = bs[:len(bs)-1]
		} else {
			for len(bs) > 0 && bs[len(bs)-1] == '1' {
				bs = bs[:len(bs)-1]
				ans++
			}
			if len(bs) == 0 {
				return ans + 1
			} else {
				bs[len(bs)-1] = '1'
			}
		}
		ans++
	}
	return ans
}

func main() {
	fmt.Println(numSteps("1101"))
	fmt.Println(numSteps("10"))
	fmt.Println(numSteps("11001"))
}
