package main

import "fmt"

func numberOfSpecialChars(word string) int {
	set := map[rune]bool{}
	for _, ch := range word {
		set[ch] = true
	}
	ans := 0
	for ch := 'A'; ch <= 'Z'; ch++ {
		lw := 'a' + ch - 'A'
		_, ok1 := set[ch]
		if !ok1 {
			continue
		}
		if _, ok2 := set[lw]; ok2 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(numberOfSpecialChars("aaAbcBC"))
	fmt.Println(numberOfSpecialChars("abc"))
	fmt.Println(numberOfSpecialChars("abBCab"))
}
