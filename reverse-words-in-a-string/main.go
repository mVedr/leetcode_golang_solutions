package main

import (
	"bytes"
	"fmt"
	"strings"
)

func solve(s string) string {
	arr := strings.Fields(s)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	var buff bytes.Buffer
	for _, ele := range arr {
		buff.WriteString(ele)
		buff.WriteString(" ")
	}
	ans := string(buff.Bytes())
	return ans[:len(ans)-1]
}

func main() {
	fmt.Println(solve("the sky is blue"))
	fmt.Println(solve("  hello world  "))
	fmt.Println(solve("a good   example"))
}
