package main

import "fmt"

func solve(num string) string {
	n, in := len(num), -1
	for i := n - 1; i >= 0; i-- {
		nn := int(num[i] - '0')
		if nn%2 == 1 {
			in = i
			break
		}
	}
	return num[:in+1]
}

func main() {
	fmt.Println(solve("52"))
	fmt.Println(solve("246"))
	fmt.Println(solve("35427"))
}
